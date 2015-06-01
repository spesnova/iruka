package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var cmdPs = cli.Command{
	Name:   "ps",
	Action: runPs,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "List containers",
	Description: `
  $ iruka ps

EXAMPLE:
  $ iruka ps
`,
}

func runPs(c *cli.Context) {
	appIdentity := getAppIdentity(c)

	containers, err := client.ContainerList(appIdentity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(containers) == 0 {
		fmt.Println("No container to list")
		os.Exit(0)
	}

	fmt.Fprintln(out, strings.Join([]string{"NAME", "IMAGE", "SIZE", "STATUS", "CREATED", "COMMAND"}, "\t"))

	for _, container := range containers {
		var f []string
		f = append(f, container.Name)
		f = append(f, container.Image)
		f = append(f, container.Size)
		f = append(f, container.State)
		f = append(f, timeDurationInWords(container.CreatedAt))
		f = append(f, container.Command)
		fmt.Fprintln(out, strings.Join(f, "\t"))
	}

	out.Flush()
}
