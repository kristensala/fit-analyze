package fit

import (
	"bytes"
	"fmt"
    "os"
	"path/filepath"

	"github.com/tormoder/fit"
)

type FitParser struct {
    File []byte
}

func (fp *FitParser) Parse() {
    file := filepath.Join("data", "tmp", "test.fit")
    fileData, err := os.ReadFile(file)
    if err != nil {
        println(err.Error())
        return
    }

    fit, err := fit.Decode(bytes.NewReader(fileData))
    if err != nil {
        println(err.Error())
        return
    }

    activity, err := fit.Activity()
    if err != nil {
        println(err.Error())
        return
    }

    for _, record := range activity.Records {
        row := fmt.Sprintf("%s, %d m, %d W, %d bpm",
            record.Timestamp.Format("2006-01-02 15:04:05"),
            record.Distance,
            record.Power,
            record.HeartRate)

        println(row)
    }
}

