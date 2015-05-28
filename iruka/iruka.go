package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/client"
)

var (
	client *iruka.Client
)

func main() {
	app := cli.NewApp()
	app.Name = "iruka"
	app.Usage = "iruka CLI"
	app.Version = "0.1.0"
	app.Author = "Seigo Uchida"
	app.Email = "spesnova@gmail.com"
	app.Commands = Commands

	client = iruka.NewClient()

	app.Run(os.Args)
}
