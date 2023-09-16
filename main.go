package main

import "github.com/kristen.sala/fit-analyze/handler"


func main() {
    summaryHandler := handler.SummaryHandler{}

    initServer(summaryHandler)
}
