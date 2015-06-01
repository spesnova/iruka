package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/spesnova/iruka/schema"
)

type Web struct {
	Command string
	Ports   []int
	Size    string
}

type Procs struct {
	App   string
	Image string
	Web   Web
}

var cmdDeploy = cli.Command{
	Name:   "deploy",
	Action: runDeploy,
	Flags:  []cli.Flag{flagAppIdentity},
	Usage:  "Deploy containers of an app",
	Description: `
  $ iruka deploy

EXAMPLE:
  $ iruka deploy
  -----> Deploying containers... done
`,
}

func runDeploy(c *cli.Context) {
	var p Procs
	err := parseYamlFile("iruka.yml", &p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if p.App == "" {
		fmt.Println("No app specified")
		fmt.Println("Run this command from an app folder or specify which app to use with --app APP.")
		os.Exit(1)
	}

	fmt.Printf("-----> Deploying containers... ")

	containers, err := client.ContainerList(p.App)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	opts := schema.ContainerCreateOpts{
		Image:   p.Image,
		Size:    p.Web.Size,
		Command: p.Web.Command,
		Type:    "web",
		Ports:   p.Web.Ports,
	}

	if len(containers) == 0 {
		err = client.ContainerCreate(p.App, opts)
	} else {
		for _, c := range containers {
			err = client.ContainerDelete(p.App, c.ID.String())
		}
		client.ContainerCreate(p.App, opts)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("done\n")
}
