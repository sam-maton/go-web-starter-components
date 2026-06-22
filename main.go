package main

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	cache     TemplateCache
	blogCache TemplateCache
}

type templateData struct {
	Pages []Metadata
}

func main() {

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	pages, err := blogPages()
	if err != nil {
		log.Fatal(err)
	}

	app := Application{
		cache: templateCache,
	}

	data := templateData{
		Pages: pages,
	}

	http.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("home.html", w, r, data)
	}))

	http.Handle("GET /blog/{slug}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		app.render(fmt.Sprintf("%s.html", slug), w, r, data)
	}))

	http.Handle("GET /components", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("components.html", w, r, data)
	}))

	http.Handle("GET /documentation", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("documentation.html", w, r, data)
	}))

	fmt.Println("Server started @ http://localhost:4321")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
