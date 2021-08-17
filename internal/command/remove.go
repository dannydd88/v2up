package command

import (
	"github.com/dannydd88/v2up/internal/user"
	"github.com/urfave/cli/v2"
)

type RemoveHolder interface {
	Remove(*cli.Context) error
}

var removeCmdMap = map[string]RemoveHolder{
	SUB_COMMAND_USER: user.NewUser(),
}

func NewRemoveCommand() *cli.Command {
	return &cli.Command{
		Name:  COMMAND_REMOVE,
		Usage: "Remove operation for specific resource",
		Subcommands: []*cli.Command{
			{
				Name:   SUB_COMMAND_USER,
				Usage:  "Remove user",
				Action: removeCmdMap[SUB_COMMAND_USER].Remove,
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowCommandHelpAndExit(c, "", 0)
			return nil
		},
	}
}
