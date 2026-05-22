package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title   string
	Heading string
}

func newHandler() (http.Handler, error) {
	tmpl, err := template.ParseFiles("templates/base.html")
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	render := func(w http.ResponseWriter, r *http.Request, data pageData) {
		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			log.Printf("template render error for path %q template %q: %v", r.URL.Path, "base", err)
			http.Error(w, "template render error", http.StatusInternalServerError)
		}
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		render(w, r, pageData{Title: "Home", Heading: "Home"})
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, pageData{Title: "About", Heading: "About"})
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
