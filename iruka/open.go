package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

var cmdOpen = cli.Command{
	Name:   "open",
	Action: runOpen,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "Open the app in a web browser",
	Description: `
  $ iruka open
`,
}

func runOpen(c *cli.Context) {
	appIdentity := getAppIdentity(c)

	containers, err := client.ContainerList(appIdentity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(containers) == 0 {
		fmt.Println("No containers are running")
		os.Exit(0)
	}

	if strings.HasPrefix(containers[0].State, "Up") == false {
		fmt.Println("No containers are running")
		os.Exit(0)
	}

	url := "http://" + containers[0].Machine + ":" + strconv.FormatInt(containers[0].PublishedPort, 10)

	fmt.Printf("Opening %s... ", appIdentity)

	_, err = exec.Command("open", url).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("done\n")
}
