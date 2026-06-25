package main

import (
	"fmt"
	"log"

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

	createPages(app, data)

	err = copyStaticFiles()
	if err != nil {
		fmt.Println(err)
	}
}
