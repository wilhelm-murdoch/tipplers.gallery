package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var spirits = []string{"absinthe", "brandy", "mezcal", "whiskey", "whisky", "rum", "gin", "vodka", "tequila", "scotch", "bourbon", "cognac"}

type Ingredient struct {
	Name        string  `json:"name"`
	Measurement float64 `json:"measurement,omitempty"`
	Unit        string  `json:"unit,omitempty"`
}

type Instruction struct {
	Description string `json:"description"`
}

type Equipment struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
}

type Tag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Image struct {
	SourceUrl    string `json:"source_url"`
	RelativePath string `json:"relative_path"`
	Attribution  string `json:"attribution,omitempty"`
}

type Meta struct {
	StandardDrinks float64 `json:"standard_drinks,omitempty"`
	SourcedFrom    string  `json:"sourced_from"`
	SourceUrl      string  `json:"source_url"`
}

type Cocktail struct {
	Cid          string        `json:"cid"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
	Images       []Image       `json:"images"`
	Tags         []Tag         `json:"tags"`
	Equipment    []Equipment   `json:"equipment"`
	GlassWare    string        `json:"glassware,omitempty"`
	Approved     bool          `json:"approved"`
	Meta         Meta          `json:"meta"`
}

type CocktailSlim struct {
	Cid    string `json:"cid"`
	Name   string `json:"n"`
	Image  string `json:"i"`
	Counts string `json:"c"`
}

func NewCocktail(name, sourcedFrom, sourceUrl string) *Cocktail {
	return &Cocktail{
		Cid:  slug.Make(fmt.Sprintf("%s-%s", sourcedFrom, name)),
		Name: name,
		Meta: Meta{
			SourcedFrom: sourcedFrom,
			SourceUrl:   sourceUrl,
		},
	}
}

func (c *Cocktail) PushTag(name string) {
	name = cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimSpace(name))

	var found = func(tag string) bool {
		for _, t := range c.Tags {
			if t.Name == tag {
				return true
			}
		}

		return false
	}

	if !found(name) {
		c.Tags = append(c.Tags, Tag{
			name, slug.Make(name),
		})
	}
}

func (c *Cocktail) PushIngredient(name string, measurement float64, unit string) {
	c.Ingredients = append(c.Ingredients, Ingredient{
		cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimSpace(name)), measurement, unit,
	})
}

func (c *Cocktail) PushInstruction(text string) {
	c.Instructions = append(c.Instructions, Instruction{
		strings.TrimSpace(text),
	})
}

func (c *Cocktail) PushEquipment(count int, name string) {
	c.Equipment = append(c.Equipment, Equipment{
		count, cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimSpace(name)), slug.Make(name),
	})
}

func (c *Cocktail) PushImage(sourceUrl, attribution string) {
	relativeUrl := md5.Sum([]byte(sourceUrl))

	c.Images = append(c.Images, Image{
		sourceUrl, fmt.Sprintf("images/cocktails/%s/%s/original.jpg", c.Cid, hex.EncodeToString(relativeUrl[:])), attribution,
	})
}

func unmarshalFromSource(sourcePath string) ([]Cocktail, error) {
	if _, err := os.Stat(sourcePath); err != nil {
		return []Cocktail{}, err
	}

	cocktailJson, err := os.ReadFile(sourcePath)
	if err != nil {
		return []Cocktail{}, err
	}

	var cocktails []Cocktail
	json.Unmarshal(cocktailJson, &cocktails)

	return cocktails, nil
}

func getDocumentFromUrl(url string) (*goquery.Document, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var backoff time.Duration

	maxAttempts := 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		response, err := client.Get(url)
		if err != nil {
			break
		}

		switch response.StatusCode {
		case 429:
			if attempt >= maxAttempts {
				log.Printf("could not get url %s after %d attempts; skipping ...", url, attempt)
				return goquery.NewDocumentFromReader(response.Body)
			}

			backoff = time.Duration(attempt) * time.Second
			log.Printf("got rate-limited on url %s; waiting another %d seconds", url, (backoff / time.Second))
			time.Sleep(backoff)
		case 200:
			return goquery.NewDocumentFromReader(response.Body)
		case 404:
			log.Printf("url %s could not be found; skipping ...", url)
			return goquery.NewDocumentFromReader(response.Body)
		}

		defer response.Body.Close()
	}

	return nil, nil
}
