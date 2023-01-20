package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Images mg.Namespace

func (Images) Download(ctx context.Context, sourcePath, writePath string) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}
	// index[0] remote path
	// index[1] local path
	var images collection.Collection[[]string]
	for _, cocktail := range cocktails {
		for _, image := range cocktail.Images {
			if err := os.MkdirAll(fmt.Sprintf("%s/%s", writePath, strings.TrimSuffix(image.RelativePath, "/original.jpg")), os.ModePerm); err != nil {
				fmt.Println(err)
			}

			images.Push([]string{image.SourceUrl, fmt.Sprintf("%s/%s", writePath, image.RelativePath)})
		}
	}

	var errorList []string

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image []string) {
		defer func() {
			bar.Add(1)
		}()

		if _, err := os.Stat(image[1]); err == nil {
			return
		}

		response, err := http.Get(image[0])
		if err != nil {
			errorList = append(errorList, fmt.Sprintf("%s: %s", image[1], err.Error()))
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			errorList = append(errorList, fmt.Sprintf("%s: %d", image[1], response.StatusCode))
			return
		}

		var data bytes.Buffer
		_, err = io.Copy(&data, response.Body)
		if err != nil {
			errorList = append(errorList, fmt.Sprintf("%s: %s", image[1], err.Error()))
			return
		}

		if err := os.WriteFile(image[1], data.Bytes(), 0644); err != nil {
			errorList = append(errorList, fmt.Sprintf("%s: %s", image[1], err.Error()))
			return
		}
	}, 10)

	if len(errorList) > 0 {
		fmt.Printf("Found the following %d errors:\n", len(errorList))
		for _, e := range errorList {
			fmt.Println(e)
		}
	}

	return nil
}
