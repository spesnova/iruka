package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/schema"
)

var cmdCreate = cli.Command{
	Name:   "create",
	Action: runCreate,
	Usage:  "Create a new app",
	Description: `
  $ iruka create <APP_NAME>

EXAMPLE:
  $ iruka create hello
  Creating hello... done
`,
}

func runCreate(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "create")
		os.Exit(1)
	}

	name := c.Args().First()
	fmt.Printf("Creating %s... ", name)

	opts := schema.AppCreateOpts{
		Name: name,
	}

	_, err := client.AppCreate(opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("done\n")

}
