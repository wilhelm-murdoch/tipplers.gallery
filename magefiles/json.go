package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
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
