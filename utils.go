package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Metadata struct {
	Title string
	Date  time.Time
}

func (app *Application) render(page string, w http.ResponseWriter, r *http.Request, data interface{}) {
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

		fmt.Println(pageMetadata)
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

		key := parts[0]
		value := parts[1]

		switch strings.ToLower(key) {
		case "title":
			metadata.Title = value
		case "date":
			date, err := time.Parse("2006-01-02", value)
			if err != nil {
				return metadata, err
			}
			metadata.Date = date
		}
	}

	return metadata, nil
}
