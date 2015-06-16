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
		cmdConfigSet,
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

var cmdConfigSet = cli.Command{
	Name:  "set",
	Usage: "Set config vars to the app",
	Description: `
   $ iruka config set <APP> <KEY1>=<VALUE1>,<KEY2>=<VALUE2>

EXAMPLE:

   $ iruka config set example HELLO=world,SOME_FLAG=1
   FOO:       bar
   BAZ:       qux
   HELLO:     world
   SOME_FLAG: 1
`,
	Action: runConfigSet,
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

func runConfigSet(c *cli.Context) {
	if len(c.Args()) != 2 {
		cli.ShowCommandHelp(c, "set")
		os.Exit(1)
	}

	appIdentity := c.Args().First()

	configVars := make(map[string]string)
	keyValuesString := c.Args()[1]
	keyValues := strings.Split(keyValuesString, ",")
	for i := range keyValues {
		keyValue := strings.Split(keyValues[i], "=")
		configVars[keyValue[0]] = keyValue[1]
	}

	res, err := client.ConfigVarsUpdate(appIdentity, configVars)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO(dtan4): restart container here

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
