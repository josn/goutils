package date

import (
    "fmt"
    "time"
)

func DateString(t Time) string {
    return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func DateUtc(date string) int64 {
    t, err := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
    if err != nil {
        return 0
    }
    return t.Local().Unix()
}
