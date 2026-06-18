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

func main() {

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := Application{
		cache: templateCache,
	}

	http.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("home.html", w, r)
	}))

	http.Handle("GET /blog/{slug}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		app.render(fmt.Sprintf("%s.html", slug), w, r)
	}))

	http.Handle("GET /components", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("components.html", w, r)
	}))

	http.Handle("GET /documentation", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render("documentation.html", w, r)
	}))

	fmt.Println("Server started @ http://localhost:4321")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
