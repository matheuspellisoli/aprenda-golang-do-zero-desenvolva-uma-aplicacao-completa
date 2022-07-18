package main

import (
	"log"
	"net/http"
	"text/template"
)

type user struct {
	Nome string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	u := user{"Jose"}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
