package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/kristen.sala/fit-analyze/internal/fit"
)

type Data struct {
     AvgPower int `json:"name"`
}

func handleServer() {
    fs := http.FileServer(http.Dir("./assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    tmpl := template.Must(template.ParseFiles("./views/index.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil)
    })

    tmplSummary := template.Must(template.ParseFiles("./templates/summary.html"))
    http.HandleFunc("/api/template/summary", func(w http.ResponseWriter, r *http.Request) {
        data := Data{
            AvgPower: 123,
        }
        tmplSummary.Execute(w, data)
    })

    http.HandleFunc("/api/fit/upload", func(w http.ResponseWriter, r *http.Request) {
        file, handler, err := r.FormFile("fitFile")
        if err != nil {
            fmt.Println(err)
            return
        }

        defer file.Close()
        println(handler.Filename)

        tempFile, err := os.CreateTemp("./data/tmp/", "upload-*.fit")
        if err != nil {
            fmt.Println(err)
            return
        }

        defer tempFile.Close()

        fileBytes, err := io.ReadAll(file)
        if err != nil {
            fmt.Println(err)
            return
        }

        tempFile.Write(fileBytes)

        if err = tempFile.Close(); err != nil {
            fmt.Println(err)
            return
        }

        decoder := fit.FitParser{
            TmpFile: *tempFile,
        }

        result, err := decoder.Parse()
        if err != nil {
            fmt.Println(err)

            http.Error(w, http.StatusText(http.StatusInternalServerError),
                http.StatusInternalServerError)
        }
        
        json.NewEncoder(w).Encode(result)
    })

    println("Listening on http://localhost:5432")
    http.ListenAndServe(":5432", nil)
}

