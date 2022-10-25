package command

import (
	"fmt"
	"io"
	"time"

	"v2up/internal"
	"v2up/internal/infra"
	"v2up/internal/v2ray"

	"github.com/codeskyblue/go-sh"
	"github.com/urfave/cli/v2"
)

func NewUserCommand() *cli.Command {
	return &cli.Command{
		Name:   internal.COMMAND_USER,
		Usage:  "User actions to v2ray vmess account",
		Before: infra.CommandInit,
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
			{
				Name:  internal.USER_COMMAND_BACKUP,
				Usage: "Backup all vmess users config to s3",
				Action: func(c *cli.Context) error {
					src := infra.GetConfig().User.Storage
					dest := fmt.Sprintf(
						"s3://%s/vpn/user-%s.json",
						infra.GetConfig().User.Backup.Bucket,
						time.Now().Format("2006_01_02_15"),
					)
					infra.GetLogger().Log("[USER-BACKUP]", "do backup:", src, "to:", dest)

					sess := sh.NewSession()
					sess.Stdout = io.Discard
					sess.Stderr = io.Discard
					err := sess.Command(
						"aws",
						"s3",
						"cp",
						src,
						dest,
					).Run()
					if err != nil {
						return err
					}

					infra.GetLogger().Log("[USER-BACKUP]", "backup success")
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			cli.ShowSubcommandHelpAndExit(ctx, 0)
			return nil
		},
	}
}
