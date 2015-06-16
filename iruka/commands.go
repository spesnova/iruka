package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	cmdApps,
	cmdCreate,
	cmdConfig,
	cmdDestroy,
	cmdPs,
	cmdDeploy,
	cmdOpen,
	cmdSet,
	cmdUnset,
}
