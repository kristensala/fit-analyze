package main

import "github.com/kristen.sala/fit-analyze/internal/fit"


func main() {
    fit_decode := fit.FitParser{}

    fit_decode.Parse()
    //handleServer()
}
