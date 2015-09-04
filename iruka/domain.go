package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"

	_ "github.com/wantedly/kujira/schema"
)

var cmdDomain = cli.Command{
	Name:  "domain",
	Usage: "Manage domains (add, list, remove)",
	Subcommands: []cli.Command{
		cmdDomainList,
		cmdDomainAdd,
		cmdDomainRemove,
	},
}

var cmdDomainList = cli.Command{
	Name:  "list",
	Usage: "List custom domains for the app",
	Description: `
   $ iruka domain list

EXAMPLE:

   $ iruka domain list
   example.com
`,
	Action: runDomainList,
}

var cmdDomainAdd = cli.Command{
	Name:  "add",
	Usage: "Add a custom domain to the app",
	Description: `
   $ iruka domain add <DOMAIN>

EXAMPLE:

   $ iruka domain add example.com
   Adding example.com to example... done
`,
	Action: runDomainAdd,
}

var cmdDomainRemove = cli.Command{
	Name:  "remove",
	Usage: "Remove a custom domain from the app",
	Description: `
   $ iruka domain remove <DOMAIN>

EXAMPLE:

   $ iruka domain remove example.com
   Removing example.com from example... done
`,
	Action: runDomainRemove,
}

func runDomainList(c *cli.Context) {
	appIdentity := getAppIdentity(c)
	domains, err := client.DomainList(appIdentity)

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range domains {
		var f []string
		f = append(f, d.Hostname)
		fmt.Fprintln(out, strings.Join(f, "\t"))
	}

	out.Flush()
}

func runDomainAdd(c *cli.Context) {

}

func runDomainRemove(c *cli.Context) {

}
