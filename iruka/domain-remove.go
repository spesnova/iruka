package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var cmdDomainRemove = cli.Command{
	Name:  "domain-remove",
	Usage: "Remove a custom domain from the app",
	Description: `
   $ iruka domain-remove <Domain>

EXAMPLE:

   $ iruka domain-remove example.com
   Removing example.com from example... done
`,
	Action: runDomainRemove,
}

func runDomainRemove(c *cli.Context) {
	appIdentity := getAppIdentity(c)
	hostname := c.Args().First()

	fmt.Printf("Removing %s from %s...", hostname, appIdentity)
	err := client.DomainDelete(appIdentity, hostname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("done")
}
