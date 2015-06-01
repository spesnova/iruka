package main

import (
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/client"
)

var (
	client *iruka.Client
	out    *tabwriter.Writer
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

	out = new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 2, '\t', 0)

	app.Run(os.Args)
}
