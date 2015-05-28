package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var cmdDestroy = cli.Command{
	Name:   "destroy",
	Action: runDestroy,
	Usage:  "Destroy an app",
	Description: `
  $ iruka destroy <APP_NAME>

EXAMPLE:
  $ iruka destroy hello
  Destroying hello... done
`,
}

func runDestroy(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "destroy")
		os.Exit(1)
	}

	name := c.Args().First()
	fmt.Printf("Destroying %s... ", name)

	err := client.AppDelete(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("done\n")
}
