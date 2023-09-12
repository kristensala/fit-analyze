package main

import (
	"net/http"
    "html/template"
)

func handleServer() {
    tmpl := template.Must(template.ParseFiles("./views/index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    println("Listening on http://localhost:5432")
    http.ListenAndServe(":5432", nil)
}
