package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/magefile/mage/mg"
)

type Bar mg.Namespace

type TheBarCocktailMedia struct {
	ContentType string `json:"contentType"`
	OptionsJson string `json:"optionsJson"`
}

type TheBarCocktailField struct {
	Alias     string                `json:"alias"`
	Text      string                `json:"text,omitempty"`
	Html      string                `json:"html,omitempty"`
	Decimal   string                `json:"decimal,omitempty"`
	TextList  []string              `json:"textList,omitempty"`
	Blocks    []TheBarCocktailBlock `json:"blocks"`
	MediaList []TheBarCocktailMedia `json:"mediaList,omitempty"`
}

type TheBarCocktailBlock struct {
	ContentType string                `json:"contentType"`
	Fields      []TheBarCocktailField `json:"fields"`
}

type TheBarCocktail struct {
	PageProps struct {
		BreadCrumbs []struct {
			ContentType        string `json:"contentType"`
			RecipeSpiritsPages []struct {
				Title string `json:"title"`
			} `json:"recipeSpiritPages"`
			ComplexityPage struct {
				Title string `json:"title"`
			} `json:"recipeComplexityPage"`
		} `json:"breadCrumbs"`
		CurrentPage struct {
			Title  string                `json:"title"`
			Fields []TheBarCocktailField `json:"fields"`
		} `json:"currentPage"`
	} `json:"pageProps"`
}

func (Bar) Scrape(ctx context.Context, sourcePath string) error {
	if _, err := os.Stat(sourcePath); err != nil {
		return fmt.Errorf("cannot locate source path %s", sourcePath)
	}

	idFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer idFile.Close()

	var searchBlocksForField = func(blockAlias, fieldAlias string, block []TheBarCocktailBlock) (TheBarCocktailField, error) {
		for _, block := range block {
			if blockAlias == block.ContentType {
				for _, field := range block.Fields {
					if field.Alias == fieldAlias {
						return field, nil
					}
				}
			}
		}

		return TheBarCocktailField{}, fmt.Errorf("alias %s not found in block %s fields", fieldAlias, blockAlias)
	}

	var searchForField = func(alias string, fields []TheBarCocktailField) (TheBarCocktailField, error) {
		for _, field := range fields {
			if field.Alias == alias {
				return field, nil
			}
		}

		return TheBarCocktailField{}, fmt.Errorf("alias %s not found in list of fields", alias)
	}

	var yieldImageFromBlocks = func(blocks []TheBarCocktailBlock) (string, error) {
		image, _ := searchBlocksForField("cmsImageItem", "file", blocks)
		if len(image.MediaList) == 0 {
			image, _ = searchBlocksForField("dchMediaItem", "file", blocks)
			if len(image.MediaList) == 0 {
				return "[MISSING]", nil
			}
		}

		var imageJson interface{}
		json.Unmarshal([]byte(image.MediaList[0].OptionsJson), &imageJson)

		domain := imageJson.(map[string]interface{})["mediaDomain"]
		key := imageJson.(map[string]interface{})["default"].(map[string]interface{})["key"]

		if domain != "" && key != "" {
			return fmt.Sprintf("%s/%s", domain, key), nil
		}

		return "", fmt.Errorf("could not yield image from blocks")
	}

	var cocktails []*Cocktail

	scanner := bufio.NewScanner(idFile)

	for scanner.Scan() {
		id := scanner.Text()

		response, err := http.Get(fmt.Sprintf("https://www.thebar.com/_next/data/y0hj7l/en-gb/recipes/%s.json", id))
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var thebar TheBarCocktail
		json.Unmarshal(body, &thebar)

		blocks, _ := searchForField("body", thebar.PageProps.CurrentPage.Fields)
		title, _ := searchBlocksForField("recipeOverview", "title", blocks.Blocks)
		if title.Text == "" {
			title, _ = searchBlocksForField("recipeHero", "title", blocks.Blocks)
		}

		var cocktail = NewCocktail(title.Text, "thebar.com", fmt.Sprintf("https://www.thebar.com/en-gb/recipes/%s", id))

		description, _ := searchBlocksForField("recipeOverview", "recipeDescription", blocks.Blocks)
		cocktail.Description = strip.StripTags(description.Html)
		if description.Html == "" {
			description, _ = searchBlocksForField("recipeHero", "description", blocks.Blocks)
			cocktail.Description = description.Text
		}

		ingredients, _ := searchBlocksForField("recipeIngredients", "recipeIngredients", blocks.Blocks)
		if len(ingredients.Blocks) == 0 {
			ingredients, _ = searchBlocksForField("ingredientsMethod", "ingredientsList", blocks.Blocks)
		}

		if len(ingredients.Blocks) == 0 {
			ingredientsWrapper, _ := searchBlocksForField("ingredientsWrapper", "ingredients", blocks.Blocks)
			ingredients, _ = searchBlocksForField("recipeIngredientList", "ingredientList", ingredientsWrapper.Blocks)
		}

		for _, ingredient := range ingredients.Blocks {
			measurement, _ := searchForField("quantity", ingredient.Fields)
			name, _ := searchForField("description", ingredient.Fields)
			unit, _ := searchForField("quantityUnits", ingredient.Fields)

			measurementAsInt, _ := strconv.Atoi(measurement.Text)

			if measurement.Text == "" || name.Text == "" || unit.Text == "" {
				measurement, _ = searchForField("quantityValue", ingredient.Fields)
				name, _ = searchForField("ingredientName", ingredient.Fields)
				unit, _ = searchForField("quantityUnit", ingredient.Fields)

				measurementAsInt, _ = strconv.Atoi(measurement.Decimal)
			}

			cocktail.PushIngredient(name.Text, measurementAsInt, unit.Text)
		}

		instructions, _ := searchBlocksForField("recipeMethod", "recipeMethod", blocks.Blocks)
		if len(instructions.Blocks) == 0 {
			instructions, _ = searchBlocksForField("ingredientsMethod", "methodListCard", blocks.Blocks)
		}

		if len(instructions.Blocks) == 0 {
			instructionsWrapper, _ := searchBlocksForField("ingredientsWrapper", "recipeMethod", blocks.Blocks)
			instructions, _ = searchBlocksForField("recipeMethod", "recipeMethod", instructionsWrapper.Blocks)
		}

		for _, instruction := range instructions.Blocks {
			if instruction.ContentType != "callToActionText" {
				title, _ := searchForField("title", instruction.Fields)
				text, _ := searchForField("text", instruction.Fields)

				cocktail.PushInstruction(title.Text, strip.StripTags(strings.Replace(text.Html, "\n", "", -1)))
			}
		}

		imageBlock, _ := searchForField("openGraphImage", thebar.PageProps.CurrentPage.Fields)
		imageUrl, _ := yieldImageFromBlocks(imageBlock.Blocks)
		if imageUrl != "" {
			cocktail.PushImage(imageUrl, "thebar.com")
		}

		equipment, _ := searchBlocksForField("recipeIngredients", "recipeEquipment", blocks.Blocks)
		var equipmentListType int
		if len(equipment.Blocks) == 0 {
			equipment, _ = searchBlocksForField("ingredientsMethod", "equipmentListCard", blocks.Blocks)
			equipmentListType = 1
		}

		if len(equipment.Blocks) == 0 {
			equipmentWrapper, _ := searchBlocksForField("ingredientsWrapper", "ingredients", blocks.Blocks)
			equipment, _ = searchBlocksForField("recipeIngredientList", "equipmentList", equipmentWrapper.Blocks)
			equipmentListType = 2
		}

		var itemCount int
		var itemText string
		for _, item := range equipment.Blocks {
			switch equipmentListType {
			case 1:
				count, _ := searchForField("quantity", item.Fields)
				name, _ := searchForField("name", item.Fields)
				countAsInt, _ := strconv.Atoi(count.Decimal)

				itemText = name.Text
				itemCount = countAsInt
			case 2:
				count, _ := searchForField("equipmentQuantity", item.Fields)
				name, _ := searchForField("equipmentName", item.Fields)
				countAsInt, _ := strconv.Atoi(count.Decimal)

				itemText = name.Text
				itemCount = countAsInt
			default:
				text, _ := searchForField("text", item.Fields)
				textParts := strings.Split(strip.StripTags(text.Html), "x")

				if len(textParts) > 1 {
					itemCount, _ = strconv.Atoi(strings.TrimSpace(textParts[0]))
					itemText = textParts[1]
				} else {
					itemCount = 1
					itemText = textParts[0]
				}
			}

			cocktail.PushEquipment(itemCount, itemText)
		}

		for _, crumb := range thebar.PageProps.BreadCrumbs {
			if crumb.ContentType == "recipePage" {
				cocktail.PushTag(crumb.ComplexityPage.Title)
				for _, spirit := range crumb.RecipeSpiritsPages {
					cocktail.PushTag(spirit.Title)
				}
			}
		}

		tags, _ := searchBlocksForField("tags", "tags", blocks.Blocks)
		for _, tag := range tags.TextList {
			cocktail.PushTag(tag)
		}

		cocktails = append(cocktails, cocktail)
	}

	if err := json.NewEncoder(os.Stdout).Encode(cocktails); err != nil {
		panic(err)
	}

	return nil
}
