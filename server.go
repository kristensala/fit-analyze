package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/kristen.sala/fit-analyze/internal/fit"
)

type Session struct {
    Name string `json:"name"`
}

func handleServer() {
    fs := http.FileServer(http.Dir("./assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    tmpl := template.Must(template.ParseFiles("./views/index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    //htmx make a post request and send .fit file
    http.HandleFunc("/api/fit/decode", func(w http.ResponseWriter, r *http.Request) {
        session := Session{
            Name: "test session",
        }
        json.NewEncoder(w).Encode(session)
    })

    http.HandleFunc("/api/fit/test-data", func(w http.ResponseWriter, r *http.Request) {
        decoder := fit.FitParser{}
        result := decoder.Parse()
        json.NewEncoder(w).Encode(result)

    })

    println("Listening on http://localhost:5432")
    http.ListenAndServe(":5432", nil)
}
