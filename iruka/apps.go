package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var cmdApps = cli.Command{
	Name:   "apps",
	Action: runApps,
	Usage:  "List apps",
	Description: `
  $ iruka apps

EXAMPLE:
  $ iruka apps
  foo
  bar
  buz
`,
}

func runApps(c *cli.Context) {
	apps, err := client.AppList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(apps) == 0 {
		fmt.Println("No apps to list")
		os.Exit(0)
	}

	for _, app := range apps {
		fmt.Println(app.Name)
	}
}
