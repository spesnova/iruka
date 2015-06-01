package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/go-yaml/yaml"
)

func getAppIdentity(c *cli.Context) string {
	var appIdentity string

	_, err := os.Stat("iruka.yml")
	if err != nil {
		appIdentity = c.String("app")
	} else {
		var p Procs
		err := parseYamlFile("iruka.yml", &p)
		if err != nil {
			fmt.Println("Failed to parse iruka.yml")
			os.Exit(1)
		}
		appIdentity = p.App
	}

	if appIdentity == "" {
		fmt.Println("No app specified")
		fmt.Println("Run this command from an app folder or specify which app to use with --app APP.")
		os.Exit(1)
	}

	return appIdentity
}

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
