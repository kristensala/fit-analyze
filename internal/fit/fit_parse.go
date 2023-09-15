package fit

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/tormoder/fit"
)

type FitParser struct {
    TmpFile os.File
}

type Record struct {
    TimeStamp time.Time `json:"timestamp"`
    Distance uint32 `json:"distance"`
    Power uint16 `json:"power"`
    HeartRate uint8 `json:"heartRate"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

type FitSummary struct {
    AvgHeartRate uint8 `json:"avgHeartRate"`
    AvgPower uint16 `json:"avgPower"`
    NormalizedPower uint16 `json:"normalizedPower"`
    AvgCadence uint8 `json:"avgCadence"`
    TotalTime uint32 `json:"totalTime"`
    TotalMovingTime uint32 `json:"totalMovingTime"`
    Distance uint32 `json:"distance"`
}

type Session struct {
    Records []Record `json:"records"`
    Summary FitSummary `json:"summary"`
}

func (fp *FitParser) Parse() (Session, error) {
    file := filepath.Join("data", "tmp", "test.fit") //use this for testing

//    file := filepath.Join("", "", fp.TmpFile.Name())
    fileData, err := os.ReadFile(file)
    if err != nil {
        println(err.Error())
        return Session{}, err
    }

    fit, err := fit.Decode(bytes.NewReader(fileData))
    if err != nil {
        println(err.Error())
        return Session{}, err
    }

    activity, err := fit.Activity()
    if err != nil {
        println(err.Error())
        return Session{}, err
    }

    // Why is this a list??
    session := activity.Sessions[0]
    summary := FitSummary{
        AvgHeartRate: session.AvgHeartRate,
        AvgPower: session.AvgPower,
        NormalizedPower: session.NormalizedPower,
        AvgCadence: session.AvgCadence,
        TotalTime: session.TotalElapsedTime,
        TotalMovingTime: session.TotalMovingTime,
        Distance: session.TotalDistance,
    }

    var records []Record
    for _, record := range activity.Records {
        lat, _ := strconv.ParseFloat(record.PositionLat.String(), 64)
        long, _ := strconv.ParseFloat(record.PositionLong.String(), 64)

        if lat == 0 || long == 0 {
            continue
        }

        records = append(records, Record{
            TimeStamp: record.Timestamp,
            Distance: record.Distance,
            Power: record.Power,
            HeartRate: record.HeartRate,
            Latitude: lat,
            Longitude: long,
        })
    }

    res := Session{
        Summary: summary,
        Records: records,
    }

    return res, nil
}

func (fp *FitParser) GetSelectedRecordsSummary([]Record) {

}
