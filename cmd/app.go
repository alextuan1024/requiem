package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
)

var app *cli.App
var version string

func init() {
	app = &cli.App{
		Name:     "requiem",
		Usage:    "requiem [command]",
		Commands: []*cli.Command{},
		Flags:    []cli.Flag{},
		Version:  version,
	}
}

func Run() error {
	return app.Run(os.Args)
}
