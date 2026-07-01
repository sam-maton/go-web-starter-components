package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type Metadata struct {
	Title    string
	Date     string
	Category string
	URL      string
	Draft    bool
}

func BlogPages() ([]Metadata, error) {
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

	slices.SortFunc(pages, func(a, b Metadata) int {
		aDate, err := time.Parse("2006-01-02", a.Date)
		if err != nil {
			return 0
		}
		bDate, err := time.Parse("2006-01-02", b.Date)
		if err != nil {
			return 0
		}
		return bDate.Compare(aDate)
	})

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

		var value string
		key := strings.TrimSpace(parts[0])

		if len(parts) > 1 {
			value = strings.TrimSpace(parts[1])
		}

		metadata.Draft = false

		switch strings.ToLower(key) {
		case "title":
			metadata.Title = value
		case "date":
			metadata.Date = value
		case "category":
			metadata.Category = value
		case "draft":
			metadata.Draft = true
		}
	}

	return metadata, nil
}
