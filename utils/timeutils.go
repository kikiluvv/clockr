package utils

import (
	"time"

	"github.com/kikiluvv/clockr/db" // needed for db.Break type
)

const TimeLayout = "15:04"
const DateLayout = "2006-01-02"

func ParseTime(t string) (time.Time, error) {
	return time.Parse(TimeLayout, t)
}

func ParseDate(d string) (time.Time, error) {
	return time.Parse(DateLayout, d)
}

func DurationInHours(start, end string) (float64, error) {
	s, err := ParseTime(start)
	if err != nil {
		return 0, err
	}
	e, err := ParseTime(end)
	if err != nil {
		return 0, err
	}
	d := e.Sub(s).Hours()
	if d < 0 {
		d += 24
	}
	return d, nil
}

func NowTimeString() string {
	return time.Now().Format(TimeLayout)
}

func TodayDateString() string {
	return time.Now().Format(DateLayout)
}

// BreakDurationHours calculates total hours from db.Break slice
func BreakDurationHours(breaks []db.Break) float64 {
	var total float64
	for _, br := range breaks {
		if br.Start != "" && br.End != "" {
			start, err1 := ParseTime(br.Start)
			end, err2 := ParseTime(br.End)
			if err1 == nil && err2 == nil {
				h := end.Sub(start).Hours()
				if h < 0 {
					h += 24
				}
				total += h
			}
		}
	}
	return total
}
