package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sam-maton/go-web-starter-components/internal/blog"
	"github.com/sam-maton/go-web-starter-components/internal/template"
)

type templateData struct {
	Pages []blog.Metadata
}

func main() {

	templateCache, err := template.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	pages, err := blog.BlogPages()
	if err != nil {
		log.Fatal(err)
	}

	app := template.Application{
		Cache: templateCache,
	}

	data := templateData{
		Pages: pages,
	}

	http.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.Render("index.html", w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	http.Handle("GET /blog/{slug}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		err := app.Render(fmt.Sprintf("blog/%s.html", slug), w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	http.Handle("GET /elements", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.Render("elements.html", w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	http.Handle("GET /components", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.Render("components.html", w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	http.Handle("GET /css-variables", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.Render("css-variables.html", w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	http.Handle("GET /documentation", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.Render("documentation.html", w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}))

	fmt.Println("Server started @ http://localhost:4321")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
