package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	cmdApps,
	cmdCreate,
	cmdDestroy,
	cmdPs,
	cmdDeploy,
	cmdOpen,
}
