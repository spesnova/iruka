package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var cmdDomains = cli.Command{
	Name:  "domains",
	Usage: "List custom domains for the app",
	Description: `
   $ iruka domains

EXAMPLE:

   $ iruka domains
   example.com
`,
	Action: runDomains,
}

func runDomains(c *cli.Context) {
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
