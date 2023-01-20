package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/magefile/mage/mg"
)

type Bevvy mg.Namespace

type BevvyCocktail struct {
	Did              string `json:"drink_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	DescriptionShort string `json:"description_short"`
	DrinkAid         string `json:"drink_aid"`
	Slug             string `json:"slug"`
	Ingredients      string `json:"ingredients"`
	Instructions     string `json:"instructions"`
	Views            string `json:"count_views"`
	ImageUrl         string `json:"image_large_url"`
	ImageCredit      string `json:"image_credit"`
}

func (Bevvy) Scrape(ctx context.Context, sourcePath string) error {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	var cocktails []*Cocktail

	var bevvy []BevvyCocktail
	json.Unmarshal(data, &bevvy)

	for _, b := range bevvy {
		viewToInt, _ := strconv.Atoi(b.Views)
		if viewToInt <= 2500 || b.ImageUrl == "" {
			continue
		}

		cocktail := NewCocktail(b.Name, "bevvy.com", fmt.Sprintf("https://bevvy.co/cocktail/%s/%s", b.Slug, b.DrinkAid))

		cocktail.Description = strip.StripTags(b.Description)
		if cocktail.Description == "" {
			cocktail.Description = strip.StripTags(b.DescriptionShort)
		}

		if strings.TrimSpace(cocktail.Description) == "" {
			continue
		}

		credit := "bevvy.com"
		if b.ImageCredit != "" {
			credit = b.ImageCredit
		}

		cocktail.PushImage(fmt.Sprintf("https:%s", b.ImageUrl), credit)

		for _, instruction := range strings.Split(b.Instructions, "|") {
			cocktail.PushInstruction(instruction)
		}

		for _, ingredient := range strings.Split(b.Ingredients, "|") {
			addedPartCount := 0
			ingredientParts := strings.Split(ingredient, " ")

			measurement := 0.0
			if m, err := strconv.ParseFloat(ingredientParts[0], 32); err == nil {
				measurement = math.Round(m*100) / 100
				addedPartCount++
			}

			unit := ""
			if len(ingredientParts) > 1 && regexp.MustCompile("(?i)"+regexp.QuoteMeta(ingredientParts[1])).MatchString("oz") {
				unit = ingredientParts[1]
				addedPartCount++
			}

			name := strings.Join(ingredientParts[addedPartCount:], " ")

			cocktail.PushIngredient(name, measurement, unit)
		}

		response, err := http.Get(fmt.Sprintf("https://api.bevvy.co/v0/tags/tags?content_id=%s&content_type=drink", b.Did))
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil
		}

		var tags []interface{}
		json.Unmarshal(body, &tags)

		for _, tag := range tags {
			cocktail.PushTag(tag.(map[string]interface{})["slug"].(string))
		}

		cocktails = append(cocktails, cocktail)
	}

	if err := json.NewEncoder(os.Stdout).Encode(cocktails); err != nil {
		panic(err)
	}

	return nil
}
