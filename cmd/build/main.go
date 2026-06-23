package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sam-maton/go-web-starter-components/internal/blog"
	"github.com/sam-maton/go-web-starter-components/internal/template"
)

type templateData struct {
	Pages []blog.Metadata
}

func main() {

	cache, err := template.NewTemplateCache()
	if err != nil {
		fmt.Println(err)
		log.Fatal("There was an error creating the template cache")
	}

	pages, err := blog.BlogPages()
	if err != nil {
		fmt.Println(err)
		log.Fatal("There was an error retrieving the blog pages")
	}

	app := template.Application{
		Cache: cache,
	}

	data := templateData{
		Pages: pages,
	}

	cleanOutputDir()

	for k := range app.Cache {
		folder := "public/" + strings.TrimSuffix(k, ".html")
		path := strings.TrimSuffix(folder, "index")
		os.MkdirAll(path, 0755)

		f, err := os.Create(path + "/index.html")
		check(err)
		defer f.Close()

		app.Render(k, f, data)
	}
}
