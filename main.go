package main

import (
	"log"
	"net/http"
)

func newHandler() (http.Handler, error) {
	templateCache, err := newTemplateCahce()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	render := func(w http.ResponseWriter, name string) {
		t, ok := templateCache[name]
		if !ok {
			http.Error(w, "template not found", http.StatusInternalServerError)
			return
		}

		if err := t.ExecuteTemplate(w, "base", nil); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
		}
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		render(w, "home.html")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about.html")
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux, nil
}

func main() {
	handler, err := newHandler()
	if err != nil {
		log.Fatalf("failed to initialize handler: %v", err)
	}

	log.Println("server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
