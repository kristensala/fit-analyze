package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type SummaryRequest struct {
    Records []RecordsRequest `json:"records"`
}

type RecordsRequest struct {
    Power int `json:"power"`
    HeartRate int `json:"heartRate"`
    Timestamp string `json:"timestamp"`
}

type SummaryResponse struct {
    AvgHeartRate int
    AvgPower int16
    Distance int64
    Duration int64 //seconds
}

type SummaryHandler struct {}

func (s SummaryHandler) HandleSummaryRequest(r *http.Request) (SummaryResponse, error) {
    var summaryRequest SummaryRequest

    b, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Error: %s", err)
        return SummaryResponse{}, err
    }

    err = json.Unmarshal(b, &summaryRequest)
    if err != nil {
        log.Fatalf("Error: %s", err)
        return SummaryResponse{}, err
    }

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

    avgHeartRate := heartRateValuesSum / heartRateValuesCount

    powerValuesCount := len(powerValues)
    var powerValuesSum int
    for _, power := range powerValues {
        powerValuesSum += power
    }
    avgPower := powerValuesSum / powerValuesCount

    //duration
    startTime := summaryRequest.Records[0].Timestamp
    endTime := summaryRequest.Records[len(summaryRequest.Records) - 1].Timestamp

    parsedStart, err := time.Parse(time.RFC3339, startTime)
    if err != nil {
        log.Fatalln(err)

    }
    parsedEnd, err := time.Parse(time.RFC3339, endTime)
    if err != nil {
        log.Fatalln(err)

    }

    duration := int64(parsedEnd.Sub(parsedStart).Seconds())

    return SummaryResponse{
        AvgHeartRate: avgHeartRate,
        AvgPower: int16(avgPower),
        Duration: duration,
    }, nil

}

