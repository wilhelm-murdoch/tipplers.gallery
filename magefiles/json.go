package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/wilhelm-murdoch/go-collection"
)

type Json mg.Namespace

func (Json) Slim(ctx context.Context, sourcePath string, approvedOnly bool) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}

	var slim []CocktailSlim
	for _, cocktail := range cocktails {
		if approvedOnly && !cocktail.Approved {
			continue
		}

		slim = append(slim, CocktailSlim{
			Cid:    cocktail.Cid,
			Name:   cocktail.Name,
			Image:  strings.TrimPrefix(cocktail.Images[0].RelativePath, fmt.Sprintf("images/cocktails/%s/", cocktail.Cid)),
			Counts: fmt.Sprintf("%d|%d|%d", len(cocktail.Ingredients), len(cocktail.Instructions), len(cocktail.Tags)),
		})
	}

	if err := json.NewEncoder(os.Stdout).Encode(slim); err != nil {
		return err
	}

	return nil
}

type TagItem struct {
	Name      string                        `json:"name"`
	Slug      string                        `json:"slug"`
	Count     int                           `json:"count"`
	Cocktails collection.Collection[string] `json:"tags"`
}

func (Json) Tags(ctx context.Context, sourcePath string, approvedOnly bool) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}

	var tags collection.Collection[*TagItem]
	for _, cocktail := range cocktails {
		for _, tag := range cocktail.Tags {
			tags.PushDistinct(&TagItem{
				Name: tag.Name,
				Slug: tag.Slug,
			})
		}
	}

	tags.Each(func(i int, t *TagItem) bool {
		for _, cocktail := range cocktails {
			for _, tag := range cocktail.Tags {
				if tag.Slug == t.Slug {
					t.Count += 1
					t.Cocktails.PushDistinct(cocktail.Cid)
				}
			}
		}
		return false
	})

	if err := json.NewEncoder(os.Stdout).Encode(tags.Items()); err != nil {
		return err
	}

	return nil
}
