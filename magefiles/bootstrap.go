package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type Bootstrap mg.Namespace

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

func (Bootstrap) Bevvy(ctx context.Context, toPath string) error {
	if _, err := os.Stat(toPath); err != nil {
		return fmt.Errorf("cannot locate destination path %s", toPath)
	}

	return nil
}
