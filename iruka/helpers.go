package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
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

func parseYamlFile(filename string, v interface{}) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("Can not read " + filename)
	}

	err = yaml.Unmarshal(buf, v)
	if err != nil {
		return errors.New("Can not parse " + filename)
	}

	return nil
}
