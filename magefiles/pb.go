package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/magefile/mage/mg"
)

type Pb mg.Namespace

type PbRecipe struct {
	Cid            string   `json:"cid"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Method         string   `json:"method"`
	ImageUrl       string   `json:"image_url"`
	SourcedUrl     string   `json:"sourced_url"`
	SourcedFrom    string   `json:"sourced_from"`
	StandardDinks  float64  `json:"standard_drinks"`
	PrimarySpirits []string `json:"primary_spirits"`
	Glassware      string   `json:"glassware"`
	IsApproved     bool     `json:"is_approved"`
}

func (Pb) Import(ctx context.Context, sourcePath, databasePath string) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Printf("Upserting %d cocktails ...\n", len(cocktails))
	for _, cocktail := range cocktails {
		var primarySpirits []string
		for _, tag := range cocktail.Tags {
			for _, spirit := range spirits {
				if regexp.MustCompile("(?i)" + regexp.QuoteMeta(spirit)).MatchString(tag.Slug) {
					primarySpirits = append(primarySpirits, spirit)
				}
			}
		}

		var methodComposite []string
		for _, method := range cocktail.Instructions {
			methodComposite = append(methodComposite, method.Description)
		}

		body, _ := json.Marshal(PbRecipe{
			Cid:            cocktail.Cid,
			Name:           cocktail.Name,
			Description:    cocktail.Description,
			Method:         strings.Join(methodComposite, "|||"),
			ImageUrl:       cocktail.Images[0].SourceUrl,
			SourcedUrl:     cocktail.Meta.SourceUrl,
			SourcedFrom:    cocktail.Meta.SourcedFrom,
			StandardDinks:  cocktail.Meta.StandardDrinks,
			IsApproved:     false,
			Glassware:      "Highball",
			PrimarySpirits: primarySpirits,
		})

		// PrimarySpirits
		// Glassware

		response, err := http.Post("http://0.0.0.0:8090/api/collections/recipes/records", "application/json", bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		var recordJson interface{}
		err = json.Unmarshal(body, &recordJson)
		if err != nil {
			panic(err)
		}

		fmt.Println(recordJson)

		primarySpirits = nil
	}

	return nil
}
