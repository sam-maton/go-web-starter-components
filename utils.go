package main

import (
	"html/template"
	"net/http"
)

func render(page string, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ui/html/base.html", "ui/html/pages/"+page, "ui/html/partials/topNav.html")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = t.Execute(w, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
