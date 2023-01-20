package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Diffords mg.Namespace

type DiffordsCocktailUrl struct {
	Slug string `json:"slug"`
	Url  string `json:"url"`
}

type DiffordsCocktail struct {
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Url            string   `json:"url"`
	Description    string   `json:"description"`
	Image          string   `json:"image"`
	Author         string   `json:"author"`
	GlassWare      string   `json:"glassWare"`
	StandardDrinks string   `json:"standardDrinks"`
	Ingredients    []string `json:"recipeIngredient"`
	Instructions   []struct {
		Description string `json:"text"`
	} `json:"recipeInstructions"`
}

func (Diffords) Scrape(ctx context.Context, sourcePath, writePath string) error {
	if _, err := os.Stat(writePath); err != nil {
		return fmt.Errorf("cannot locate write path %s", writePath)
	}

	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var diffords []DiffordsCocktailUrl
	json.Unmarshal(data, &diffords)

	urls := collection.New(diffords...)

	var standardDrinks string
	var glassWare string

	bar := progressbar.Default(int64(urls.Length()))
	urls.Batch(func(b, j int, dc DiffordsCocktailUrl) {
		defer func() {
			bar.Add(1)
		}()

		if _, err := os.Stat(fmt.Sprintf("%s/%s.json", writePath, dc.Slug)); err == nil {
			return
		}

		document, err := getDocumentFromUrl(dc.Url)
		if err != nil {
			panic(err)
		}

		document.Find("article a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			if strings.Contains(href, "cocktail-glassware") {
				glassWare = s.Text()
			}
		})

		document.Find("article li").Each(func(i int, s *goquery.Selection) {
			if strings.Contains(s.Text(), "standard drinks") {
				standardDrinks = strings.TrimSpace(strings.Replace(s.Text(), "standard drinks", "", -1))
			}
		})

		document.Find("script").Each(func(i int, s *goquery.Selection) {
			if t, ok := s.Attr("type"); t == "application/ld+json" && ok {
				var j interface{}
				json.Unmarshal([]byte(s.Text()), &j)

				recipeMap, _ := j.(map[string]interface{})

				recipeMap["standardDrinks"] = standardDrinks
				recipeMap["glassWare"] = glassWare
				recipeMap["slug"] = dc.Slug
				recipeMap["url"] = dc.Url

				saveJson, _ := json.Marshal(j)

				if recipeMap["@type"] == "Recipe" && recipeMap["recipeCategory"] == "Drink" && recipeMap["recipeCuisine"] == "Cocktail" {
					if err := os.WriteFile(fmt.Sprintf("%s/%s.json", writePath, dc.Slug), saveJson, 0644); err != nil {
						panic(err)
					}
				}
			}
		})
	}, 5)

	return nil
}

func (Diffords) Munge(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	var fileList collection.Collection[string]

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			fileList.Push(fmt.Sprintf("%s/%s", sourcePath, file.Name()))
		}
	}

	var cocktails []*Cocktail

	fileList.Batch(func(i1, i2 int, s string) {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(err)
		}

		var diffords DiffordsCocktail
		json.Unmarshal(data, &diffords)

		cocktail := NewCocktail(diffords.Name, "diffordsguide.com", diffords.Url)

		cocktail.Description = diffords.Description
		cocktail.GlassWare = cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimSpace(diffords.GlassWare))

		switch count := len(diffords.Instructions); {
		case count <= 2:
			cocktail.PushTag("Beginner")
		case count > 2 && count <= 4:
			cocktail.PushTag("Intermediate")
		case count > 4:
			cocktail.PushTag("Advanced")
		}

		for _, instruction := range diffords.Instructions {
			cocktail.PushInstruction(instruction.Description)
		}

		for _, ingredient := range diffords.Ingredients {
			addedPartCount := 0
			ingredientParts := strings.Split(ingredient, " ")

			measurement := 0.0
			if m, err := strconv.ParseFloat(ingredientParts[0], 32); err == nil {
				measurement = math.Round(m*100) / 100
				addedPartCount++
			}

			unit := ""
			if ingredientParts[1] == "ml" {
				unit = ingredientParts[1]
				addedPartCount++
			}

			name := strings.Join(ingredientParts[addedPartCount:], " ")

			for _, spirit := range spirits {
				if regexp.MustCompile("(?i)" + regexp.QuoteMeta(spirit)).MatchString(name) {
					cocktail.PushTag(spirit)
				}
			}

			cocktail.PushIngredient(name, measurement, unit)
		}

		cocktail.PushTag(diffords.GlassWare)
		cocktail.PushImage(diffords.Image, "Difford's Guide")

		unitsToFloat, _ := strconv.ParseFloat(diffords.StandardDrinks, 32)
		cocktail.Meta.StandardDrinks = math.Round(unitsToFloat*100) / 100

		cocktails = append(cocktails, cocktail)
	}, 20)

	if err := json.NewEncoder(os.Stdout).Encode(cocktails); err != nil {
		panic(err)
	}

	return nil
}
