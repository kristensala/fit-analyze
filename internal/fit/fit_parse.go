package fit

import (
	"bytes"
	"os"
	"path/filepath"
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
}

func (fp *FitParser) Parse() []Record {
    file := filepath.Join("data", "tmp", "test.fit")
    fileData, err := os.ReadFile(file)
    if err != nil {
        println(err.Error())
        return nil
    }

    fit, err := fit.Decode(bytes.NewReader(fileData))
    if err != nil {
        println(err.Error())
        return nil
    }

    activity, err := fit.Activity()
    if err != nil {
        println(err.Error())
        return nil
    }

    var result []Record
    for _, record := range activity.Records {
        result = append(result, Record{
            TimeStamp: record.Timestamp,
            Distance: record.Distance,
            Power: record.Power,
            HeartRate: record.HeartRate,
        })


    }

    return result
}

