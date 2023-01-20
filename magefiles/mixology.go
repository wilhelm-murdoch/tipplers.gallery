package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Mixology mg.Namespace

type MixologyCocktail struct {
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Images      []string `json:"image"`
	Author      struct {
		Name string `json:"name"`
	} `json:"author"`
	GlassWare      string   `json:"glassWare"`
	StandardDrinks string   `json:"standardDrinks"`
	Ingredients    []string `json:"recipeIngredient"`
	Instructions   []struct {
		Description string `json:"text"`
		Url         string `json:"url"`
	} `json:"recipeInstructions"`
}

func (Mixology) Collect(ctx context.Context, spiritName string) error {
	var makeRange = func(min, max int) []int {
		a := make([]int, max-min+1)
		for i := range a {
			a[i] = min + i
		}
		return a
	}

	var (
		urls    collection.Collection[string]
		baseUrl = fmt.Sprintf("https://advancedmixology.com/blogs/art-of-mixology/tagged/%s-cocktails", spiritName)
		pages   *collection.Collection[int]
	)

	document, err := getDocumentFromUrl(baseUrl)
	if err != nil {
		return err
	}

	document.Find("ul.pagination-custom li:nth-child(6) a").Each(func(i int, s *goquery.Selection) {
		totalPages, _ := strconv.Atoi(s.Text())
		pageRange := makeRange(1, totalPages)
		pages = collection.New(pageRange...)
	})

	pages.Batch(func(i1, i2, p int) {
		document, err = getDocumentFromUrl(fmt.Sprintf("%s?page=%d", baseUrl, p))
		if err != nil {
			panic(err)
		}

		document.Find(".shopify-section h2 a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			if strings.Contains(href, "-recipe") {
				urls.PushDistinct(href)
			}
		})
	}, 5)

	if err := json.NewEncoder(os.Stdout).Encode(urls.Items()); err != nil {
		return err
	}

	return nil
}

func (Mixology) Scrape(ctx context.Context, sourcePath, writePath string) error {
	if _, err := os.Stat(writePath); err != nil {
		return fmt.Errorf("cannot locate write path %s", writePath)
	}

	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var recipeUrls []string
	json.Unmarshal(data, &recipeUrls)

	urls := collection.New(recipeUrls...)

	bar := progressbar.Default(int64(urls.Length()))
	urls.Batch(func(b, j int, url string) {
		defer func() {
			bar.Add(1)
		}()

		urlParts := strings.Split(url, "/")
		slug := urlParts[len(urlParts)-1]
		if _, err := os.Stat(fmt.Sprintf("%s/%s.json", writePath, slug)); err == nil {
			return
		}

		recipePage, err := getDocumentFromUrl(fmt.Sprintf("https://advancedmixology.com%s", url))
		if err != nil {
			panic(err)
		}

		recipePage.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
			if t, ok := s.Attr("type"); t == "application/ld+json" && ok {
				var j interface{}
				json.Unmarshal([]byte(s.Text()), &j)

				recipeMap, _ := j.(map[string]interface{})
				saveJson, _ := json.Marshal(j)

				if recipeMap["@type"] == "Recipe" {
					if err := os.WriteFile(fmt.Sprintf("%s/%s.json", writePath, slug), saveJson, 0644); err != nil {
						panic(err)
					}
				}
			}
		})
	}, 10)

	return nil
}

func (Mixology) Munge(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		return err
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

		var mixology MixologyCocktail
		json.Unmarshal(data, &mixology)

		cocktail := NewCocktail(mixology.Name, "advancedmixology.com", "")

		cocktail.Description = mixology.Description

		switch count := len(mixology.Instructions); {
		case count <= 2:
			cocktail.PushTag("Beginner")
		case count > 2 && count <= 4:
			cocktail.PushTag("Intermediate")
		case count > 4:
			cocktail.PushTag("Advanced")
		}

		for _, instruction := range mixology.Instructions {
			cocktail.PushInstruction(instruction.Description)
		}

		for _, ingredient := range mixology.Ingredients {
			for _, spirit := range spirits {
				if regexp.MustCompile("(?i)" + regexp.QuoteMeta(spirit)).MatchString(ingredient) {
					cocktail.PushTag(spirit)
				}
			}

			cocktail.PushIngredient(ingredient, 0, "")
		}

		cocktail.PushImage(mixology.Images[len(mixology.Images)-1], mixology.Author.Name)

		cocktail.Meta.SourceUrl = strings.TrimSuffix(mixology.Instructions[0].Url, "#step-1")

		cocktails = append(cocktails, cocktail)
	}, 20)

	if err := json.NewEncoder(os.Stdout).Encode(cocktails); err != nil {
		panic(err)
	}

	return nil
}
