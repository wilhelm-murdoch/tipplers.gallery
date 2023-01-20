package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/schollz/progressbar/v3"
	"github.com/wilhelm-murdoch/go-collection"
)

type Cloudflare mg.Namespace

type CloudflareImage struct {
	Url, Id string
}

func (Cloudflare) Upload(ctx context.Context, sourcePath string, approvedOnly bool) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}

	var images collection.Collection[CloudflareImage]
	for _, cocktail := range cocktails {
		if approvedOnly && !cocktail.Approved {
			continue
		}

		for _, image := range cocktail.Images {
			pathParts := strings.Split(image.RelativePath, "/")
			images.Push(CloudflareImage{image.SourceUrl, strings.Join(pathParts[:len(pathParts)-1], "/")})
		}
	}

	client := &http.Client{}

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image CloudflareImage) {
		defer func() {
			bar.Add(1)
		}()

		var data bytes.Buffer
		var fw io.Writer
		form := multipart.NewWriter(&data)

		fw, err := form.CreateFormField("url")
		if err != nil {
			panic(err)
		}
		fw.Write([]byte(image.Url))

		fw, err = form.CreateFormField("id")
		if err != nil {
			panic(err)
		}
		fw.Write([]byte(image.Id))

		form.Close()

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/images/v1", os.Getenv("CLOUDFLARE_ACCOUNT_ID")), &data)
		if err != nil {
			panic(err)
		}

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_API_TOKEN")))
		request.Header.Set("Content-Type", form.FormDataContentType())

		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if response.StatusCode != 200 {
			fmt.Printf("Failed: %s\n%s\n\n", image.Id, string(body))
		}

		time.Sleep(1 * time.Second)
	}, 25)

	return nil
}

func (Cloudflare) Purge(ctx context.Context, sourcePath string) error {
	cocktails, err := unmarshalFromSource(sourcePath)
	if err != nil {
		return err
	}

	var images collection.Collection[CloudflareImage]
	for _, cocktail := range cocktails {
		for _, image := range cocktail.Images {
			images.Push(CloudflareImage{image.SourceUrl, fmt.Sprintf("images/%s", strings.TrimSuffix(image.RelativePath, "/large.jpg"))})
		}
	}

	client := &http.Client{}

	bar := progressbar.Default(int64(images.Length()))
	images.Batch(func(b, j int, image CloudflareImage) {
		defer func() {
			bar.Add(1)
		}()

		request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/images/v1/%s", os.Getenv("CLOUDFLARE_ACCOUNT_ID"), image.Id), nil)
		if err != nil {
			panic(err)
		}

		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_API_TOKEN")))

		response, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		time.Sleep(1 * time.Second)
	}, 25)

	return nil
}
