package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gosimple/slug"
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
			SpiritsPages []struct {
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

	response, err := http.Get(fmt.Sprintf("https://www.thebar.com/_next/data/y0hj7l/en-gb/recipes/%s.json", "gordons-alcohol-free-spirit-winter-tonic"))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	id := "gordons-alcohol-free-spirit-winter-tonic"

	var thebar TheBarCocktail
	json.Unmarshal(body, &thebar)

	var cocktail = &Cocktail{
		Cid:         slug.Make(thebar.PageProps.CurrentPage.Title),
		Name:        thebar.PageProps.CurrentPage.Title,
		SourcedFrom: "thebar.com",
		SourceUrl:   fmt.Sprintf("https://www.thebar.com/_next/data/y0hj7l/en-gb/recipes/%s.json", id),
	}

	if err := json.NewEncoder(os.Stdout).Encode(cocktail); err != nil {
		panic(err)
	}

	// for _, cocktail := range thebar.Items {
	// 	fmt.Println(cocktail.Description)
	// }

	return nil
}
