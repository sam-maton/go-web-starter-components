package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	cache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cache.render("home.html", w, r)
	}))

	http.Handle("/documentation", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cache.render("documentation.html", w, r)
	}))

	fmt.Println("Server started @ http://localhost:4321")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
