package main

import (
	"strconv"
	"time"
)

func timeDurationInWords(t time.Time) string {
	duration := time.Now().Sub(t)

	var d float64
	var unit string

	switch {
	case duration < 1*time.Minute:
		d = duration.Seconds()
		unit = "seconds"
	case duration < 1*time.Hour:
		d = duration.Minutes()
		unit = "minutes"
	default:
		d = duration.Hours()
		unit = "hours"
	}
	return strconv.Itoa(int(d)) + " " + unit
}
