package main

import (
	"fmt"
	"log"
	"os"

	"v2up/internal/command"
	"v2up/internal/infra"

	"github.com/urfave/cli/v2"
)

var (
	version string = "dev"
	build   string = "dev"
	sha     string = "dev"
)

func main() {
	app := &cli.App{
		Name:    "v2up",
		Version: fmt.Sprintf("%s-%s-%s", version, build, sha),
		Usage:   "control v2ray process via grpc interface",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
				Usage:   "Load config from yaml file",
			},
		},
		Before: infra.AppInit,
		Commands: []*cli.Command{
			command.NewUserCommand(),
			command.NewProcessCommand(),
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelpAndExit(c, 0)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
