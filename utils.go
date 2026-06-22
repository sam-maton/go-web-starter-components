package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Metadata struct {
	Title string
	Date  string
	URL   string
}

func (app *Application) render(page string, w http.ResponseWriter, r *http.Request, data any) {
	t, ok := app.cache[page]
	if !ok {
		http.Error(w, "Page not in cache", 500)
	}

	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cannot execute template", 500)
	}

}

func blogPages() ([]Metadata, error) {
	pages := []Metadata{}

	paths, err := filepath.Glob("./ui/html/blog/*.html")
	if err != nil {
		return nil, err
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		pageMetadata, err := parseMetadata(string(data))

		before, _ := strings.CutSuffix(filepath.Base(path), ".html")

		pageMetadata.URL = before
		pages = append(pages, pageMetadata)
	}

	return pages, nil
}

func parseMetadata(content string) (Metadata, error) {

	var metadata Metadata

	metaStart := strings.Index(content, "<!--")
	metaEnd := strings.Index(content, "-->")

	block := strings.TrimSpace(content[metaStart+4 : metaEnd])

	for l := range strings.SplitSeq(block, "\n") {
		l = strings.TrimSpace(l)

		parts := strings.SplitN(l, ":", 2)

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch strings.ToLower(key) {
		case "title":
			metadata.Title = value
		case "date":
			metadata.Date = value
		}
	}

	return metadata, nil
}
