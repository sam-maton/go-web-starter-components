package main

import (
	"log"
	"net/http"
)

func (app *Application) render(page string, w http.ResponseWriter, r *http.Request) {
	t, ok := app.cache[page]
	if !ok {
		http.Error(w, "Page not in cache", 500)
	}

	err := t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Cannot execute template", 500)
	}

}
