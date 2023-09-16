package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SummaryRequest struct {
    Records []RecordsRequest `json:"records"`
}

type RecordsRequest struct {
    Power int `json:"power"`
    HeartRate int `json:"heartRate"`
}

type SummaryHandler struct {}

func (s SummaryHandler) HandleSummaryRequest(r *http.Request) {
    var summaryRequest SummaryRequest

    b, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }

    err = json.Unmarshal(b, &summaryRequest)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }

    // TODO: calculate summary
    var heartRateValues []int
    var powerValues []int
    for _, record := range summaryRequest.Records {
        heartRateValues = append(heartRateValues, record.HeartRate)
        powerValues = append(powerValues, record.Power)
    }

    heartRateValuesCount := len(heartRateValues)
    var heartRateValuesSum int
    for _, heartRate := range heartRateValues {
        heartRateValuesSum += heartRate
    }

    println("avg heart rate", heartRateValuesSum / heartRateValuesCount)
}

