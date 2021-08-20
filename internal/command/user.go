package command

import (
	"github.com/dannydd88/v2up/internal"
	"github.com/dannydd88/v2up/internal/v2ray"
	"github.com/urfave/cli/v2"
)

func NewUserCommand() *cli.Command {
	return &cli.Command{
		Name:  internal.COMMAND_USER,
		Usage: "User actions to v2ray vmess account",
		Subcommands: []*cli.Command{
			{
				Name:   internal.USER_COMMAND_ADD,
				Usage:  "Add vmess user to v2ray",
				Action: v2ray.UserHandler().Add,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     internal.FLAG_USER_EMAIL,
						Usage:    "User email",
						Required: true,
					},
					&cli.StringFlag{
						Name:  internal.FLAG_USER_UUID,
						Usage: "User uuid",
					},
					&cli.BoolFlag{
						Name:  internal.FLAG_USER_SILENT,
						Usage: "Do not notify user via email when add user success",
						Value: false,
					},
				},
			},
			{
				Name:   internal.USER_COMMAND_REMOVE,
				Usage:  "Remove vmess user to v2ray",
				Action: v2ray.UserHandler().Remove,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     internal.FLAG_USER_EMAIL,
						Usage:    "User email",
						Required: true,
					},
				},
			},
			{
				Name:   internal.USER_COMMAND_RESTORE,
				Usage:  "Restore all vmess users to v2ray process",
				Action: v2ray.UserHandler().Restore,
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowCommandHelpAndExit(c, "", 0)
			return nil
		},
	}
}
