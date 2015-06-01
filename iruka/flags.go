package main

import (
	"github.com/codegangsta/cli"
)

var flagAppIdentity = cli.StringFlag{
	Name:  "app",
	Value: "",
	Usage: "app name or id",
}
