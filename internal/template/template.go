package template

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func NewTemplateCache() (TemplateCache, error) {
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

	blogs, err := filepath.Glob("./ui/html/blog/*.html")
	if err != nil {
		return nil, err
	}

	for _, blog := range blogs {
		name := filepath.Base(blog)

		t, err := template.ParseFiles("./ui/html/base.html", "./ui/html/partials/topbar.html")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseFiles(blog)
		if err != nil {
			return nil, err
		}

		cache[name] = t
	}

	t, err := template.ParseFiles("./ui/html/base.html", "./ui/html/home.html", "./ui/html/partials/topbar.html")
	if err != nil {
		return nil, err
	}

	cache["home.html"] = t

	return cache, nil
}
