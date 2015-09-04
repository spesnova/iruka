package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/schema"
)

var cmdDomainAdd = cli.Command{
	Name:  "domain-add",
	Usage: "Add a custom domain to the app",
	Description: `
   $ iruka domain- add <DOMAIN>

EXAMPLE:

   $ iruka domain-add example.com
   Adding example.com to example... done
`,
	Action: runDomainAdd,
}

func runDomainAdd(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, "add")
		os.Exit(1)
	}

	appIdentity := getAppIdentity(c)
	hostname := c.Args().First()
	opts := schema.DomainCreateOpts{
		Hostname: hostname,
	}

	fmt.Printf("Adding %s to %s...", hostname, appIdentity)
	_, err := client.DomainCreate(appIdentity, opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("done")
}

func runDomainRemove(c *cli.Context) {

}
