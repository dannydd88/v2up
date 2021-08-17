package command

import (
	"github.com/dannydd88/v2up/internal/user"
	"github.com/urfave/cli/v2"
)

type AddHolder interface {
	Add(*cli.Context) error
}

var addCmdMap = map[string]AddHolder{
	SUB_COMMAND_USER: user.NewUser(),
}

func NewAddCommand() *cli.Command {
	return &cli.Command{
		Name:  COMMAND_ADD,
		Usage: "Add operation for specific resource",
		Subcommands: []*cli.Command{
			{
				Name:   SUB_COMMAND_USER,
				Usage:  "Add user",
				Action: addCmdMap[SUB_COMMAND_USER].Add,
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowCommandHelpAndExit(c, "", 0)
			return nil
		},
	}
}
