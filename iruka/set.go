package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var cmdSet = cli.Command{
	Name:   "set",
	Action: runSet,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "Set config vars to the app",
	Description: `
   $ iruka set <KEY1>=<VALUE1>,<KEY2>=<VALUE2>

EXAMPLE:

   $ iruka set HELLO=world,SOME_FLAG=1
   FOO:       bar
   BAZ:       qux
   HELLO:     world
   SOME_FLAG: 1
`,
}

func runSet(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "set")
		os.Exit(1)
	}

	appIdentity := getAppIdentity(c)

	configVars := make(map[string]string)
	keyValuesString := c.Args().First()
	keyValues := strings.Split(keyValuesString, ",")
	for i := range keyValues {
		keyValue := strings.Split(keyValues[i], "=")

		if len(keyValue) != 2 {
			cli.ShowCommandHelp(c, "set")
			os.Exit(1)
		}

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
