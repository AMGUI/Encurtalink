package net

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func ServerLocal() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templates.ExecuteTemplate(w, "home.html", nil)

	})

	http.ListenAndServe(":8081", nil)

}
