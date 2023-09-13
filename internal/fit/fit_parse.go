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
    FilePath string
}

type Record struct {
    TimeStamp time.Time `json:"timestamp"`
    Distance uint32 `json:"distance"`
    Power uint16 `json:"power"`
    HeartRate uint8 `json:"heartRate"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

func (fp *FitParser) Parse() ([]Record, error) {
    file := filepath.Join("data", "tmp", "test.fit")
    fileData, err := os.ReadFile(file)
    if err != nil {
        println(err.Error())
        return nil, err
    }

    fit, err := fit.Decode(bytes.NewReader(fileData))
    if err != nil {
        println(err.Error())
        return nil, err
    }

    activity, err := fit.Activity()
    if err != nil {
        println(err.Error())
        return nil, err
    }

    var result []Record
    for _, record := range activity.Records {
        lat, _ := strconv.ParseFloat(record.PositionLat.String(), 64)
        long, _ := strconv.ParseFloat(record.PositionLong.String(), 64)
        result = append(result, Record{
            TimeStamp: record.Timestamp,
            Distance: record.Distance,
            Power: record.Power,
            HeartRate: record.HeartRate,
            Latitude: lat,
            Longitude: long,
        })
    }

    return result, nil
}

