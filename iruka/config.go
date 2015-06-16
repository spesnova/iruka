package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var cmdConfig = cli.Command{
	Name:   "config",
	Action: runConfig,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "Display config vars for the app",
	Description: `
   $ iruka config

EXAMPLE:

   $ iruka config
   FOO: bar
   BAZ: qux
`,
}

func runConfig(c *cli.Context) {
	appIdentity := getAppIdentity(c)

	res, err := client.ConfigVarsInfo(appIdentity)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		os.Exit(1)
	}

	out = showConfigVars(out, res)
	out.Flush()
}
