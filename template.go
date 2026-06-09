package main

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func newTemplateCache() (TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		t, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = t
	}

	return cache, nil
}
