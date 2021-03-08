package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	libgen "github.com/irth/golibgen"
)

func DownloadBook(b libgen.Book, downloadLocation string) error {
	err := os.MkdirAll(downloadLocation, 0755)
	if err != nil {
		return fmt.Errorf("couldn't create download directory: %w", err)
	}

	url, err := b.DownloadLink()
	if err != nil {
		return fmt.Errorf("couldn't get download link: %w", err)
	}

	filename := path.Base(url)

	println("Downloading", url, "to", filename)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("while requesting the file: %w", err)
	}
	defer resp.Body.Close()
	println("get request got")

	path := filepath.Join(downloadLocation, filename)
	println("saving to", path)
	// TODO: actually make folders for author etc if possible to determine
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("while creating the file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("while writing the book file: %w", err)
	}
	return nil
}
