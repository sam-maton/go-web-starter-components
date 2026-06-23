package template

import (
	"errors"
	"html/template"
	"io"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

type Application struct {
	Cache TemplateCache
}

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
		name := "blog/" + filepath.Base(blog)

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

	t, err := template.ParseFiles("./ui/html/base.html", "./ui/html/index.html", "./ui/html/partials/topbar.html")
	if err != nil {
		return nil, err
	}

	cache["index.html"] = t

	return cache, nil
}

func (app *Application) Render(page string, w io.Writer, data any) error {
	t, ok := app.Cache[page]
	if !ok {
		return errors.New("Page not in cache")
	}

	err := t.Execute(w, data)
	if err != nil {
		return errors.New("Cannot execute template")
	}

	return nil
}
