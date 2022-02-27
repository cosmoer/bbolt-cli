package app

import (
	"fmt"
	"runtime"

	"github.com/cosmoer/bbolt-cli/commands/dump"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	// Package is filled at linking time
	Package = "github.com/cosmoer/bbolt-cli"

	// Version holds the complete version number. Filled in at linking time.
	Version = "0.1.0"

	// Revision is filled with the VCS (e.g. git) revision being used to build
	// the program at linking time.
	Revision = ""

	// GoVersion is Go tree's version.
	GoVersion = runtime.Version()
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Name, Package, c.App.Version)
	}
}

// New returns a *cli.App instance.
func New() *cli.App {
	app := cli.NewApp()
	app.Name = "bbolt-cli"
	app.Version = Version
	app.Description = `bbolt-cli is a tool for boltDB.`
	app.Usage = `bbolt-cli`
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in logs",
		},
	}
	app.Commands = append([]cli.Command{dump.Command})
	app.Before = func(context *cli.Context) error {
		if context.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}
	return app
}
