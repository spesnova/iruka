package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/schema"
)

var cmdConfig = cli.Command{
	Name:  "config",
	Usage: "manage config vars for an app(set, show, unset)",
	Subcommands: []cli.Command{
		cmdConfigShow,
	},
}

var cmdConfigShow = cli.Command{
	Name:  "show",
	Usage: "Display config vars for the app",
	Description: `
   $ iruka config show <APP>

EXAMPLE:

   $ iruka config show example
   FOO: bar
   BAZ: qux
`,
	Action: runConfigShow,
}

func runConfigShow(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "show")
		os.Exit(1)
	}

	appIdentity := c.Args().First()

	res, err := client.ConfigVarsList(appIdentity)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		os.Exit(1)
	}

	out = showConfigVars(out, res)
	out.Flush()
}

func showConfigVars(out *tabwriter.Writer, c schema.ConfigVars) *tabwriter.Writer {
	for k, v := range c {
		var f []string
		f = append(f, k+":")
		f = append(f, v)
		fmt.Fprintln(out, strings.Join(f, "\t"))
	}
	return out
}
