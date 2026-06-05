package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("ui/html/pages/home.html")
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		err = t.Execute(w, nil)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}))

	fmt.Println("Server started @ http://localhost:4321")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
