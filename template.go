package main

import (
	"html/template"
	"net/http"
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

func (c TemplateCache) render(page string, w http.ResponseWriter, r *http.Request) {
	t, ok := c[page]
	if !ok {
		http.Error(w, "Cannot load page", 500)
	}

	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Cannot load page", 500)
	}

}
