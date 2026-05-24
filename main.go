package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
)

type pageData struct {
	Title string
}

type templateCache map[string]*template.Template

func newTemplateCache() (templateCache, error) {
	pages := []string{"home", "about"}
	cache := make(templateCache)

	for _, page := range pages {
		tmpl, err := template.ParseFiles(
			filepath.Join("templates", "base.html"),
			filepath.Join("templates", page+".html"),
		)
		if err != nil {
			return nil, err
		}
		cache[page] = tmpl
	}

	return cache, nil
}

func (c templateCache) names() []string {
	names := make([]string, 0, len(c))
	for name := range c {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func newHandler() (http.Handler, error) {
	cache, err := newTemplateCache()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	render := func(w http.ResponseWriter, r *http.Request, name string, data pageData) {
		tmpl, ok := cache[name]
		if !ok {
			log.Printf("template missing from cache for path %q template %q - available templates: %s", r.URL.Path, name, strings.Join(cache.names(), ", "))
			http.Error(w, "template render error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			log.Printf("template render error for path %q template %q: %v", r.URL.Path, name, err)
			http.Error(w, "template render error", http.StatusInternalServerError)
		}
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		render(w, r, "home", pageData{Title: "Home"})
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "about", pageData{Title: "About"})
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux, nil
}

func main() {
	handler, err := newHandler()
	if err != nil {
		log.Fatalf("failed to initialize handler: %v", err)
	}

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
