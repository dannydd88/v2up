package main

import (
	"log"
	"os"

	"github.com/dannydd88/v2up/internal/command"
	"github.com/dannydd88/v2up/internal/infra"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "v2up",
		Usage: "control v2ray process via grpc interface",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
				Usage:   "Load config from yaml file",
			},
		},
		Before: infra.Init,
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
