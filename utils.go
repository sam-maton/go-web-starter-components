package main

import "net/http"

func (app *Application) render(page string, w http.ResponseWriter, r *http.Request) {
	t, ok := app.cache[page]
	if !ok {
		http.Error(w, "Cannot load page", 500)
	}

	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Cannot load page", 500)
	}

}
