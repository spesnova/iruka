package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var cmdUnset = cli.Command{
	Name:   "unset",
	Action: runUnset,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "Remove config vars from app",
	Description: `
   $ iruka unset <KEY1>,<KEY2>

EXAMPLE:

   $ iruka unset example FOO,BAZ
   HELLO:     world
   SOME_FLAG: 1
`,
}

func runUnset(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "unset")
		os.Exit(1)
	}

	appIdentity := getAppIdentity(c)

	configVars := make(map[string]string)
	keysString := c.Args().First()
	keys := strings.Split(keysString, ",")
	for i := range keys {
		configVars[keys[i]] = ""
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
