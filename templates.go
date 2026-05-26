package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCahce() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles("./templates/base.html", page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil

}
