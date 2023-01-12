package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Ingredient struct {
	Name        string `json:"name"`
	Measurement int    `json:"measurement,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

type Instruction struct {
	Order       int    `json:"order"`
	Title       string `json:"title"`
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

type Cocktail struct {
	Cid          string        `json:"cid"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
	Images       []Image       `json:"images"`
	Tags         []Tag         `json:"tags"`
	Equipment    []Equipment   `json:"equipment"`
	GlassType    string        `json:"glass_type"`
	SourcedFrom  string        `json:"sourced_from"`
	SourceUrl    string        `json:"source_url"`
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
