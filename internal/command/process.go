package command

import (
	"io/ioutil"
	"time"

	"v2up/internal"
	"v2up/internal/infra"
	"v2up/internal/v2ray"

	"github.com/codeskyblue/go-sh"
	"github.com/urfave/cli/v2"
)

func NewProcessCommand() *cli.Command {
	return &cli.Command{
		Name:  internal.COMMAND_PROCESS,
		Usage: "Process actions to v2ray process",
		Subcommands: []*cli.Command{
			{
				Name:    internal.PROCESS_COMMAND_START,
				Usage:   "Start v2ray process",
				Aliases: []string{internal.PROCESS_COMMAND_RESTART},
				Action: func(c *cli.Context) error {
					infra.GetLogger().Log("[PROC]", "do start v2ray...")
					sess := sh.NewSession()
					sess.Stdout = ioutil.Discard
					sess.Stderr = ioutil.Discard
					err := sess.Command("systemctl", "restart", "v2ray").Run()
					if err != nil {
						return err
					}
					infra.GetLogger().Log("[PROC]", "start v2ray success, wait for 10s for process ready")
					<-time.After(10 * time.Second)
					err = v2ray.UserHandler().Restore(c)
					if err != nil {
						return err
					}
					infra.GetLogger().Log("[PROC]", "restore user to v2ray success")
					return nil
				},
			},
			{
				Name:  internal.PROCESS_COMMAND_STOP,
				Usage: "Stop v2ray process",
				Action: func(c *cli.Context) error {
					infra.GetLogger().Log("[PROC]", "do stop v2ray...")
					sess := sh.NewSession()
					sess.Stdout = ioutil.Discard
					sess.Stderr = ioutil.Discard
					err := sess.Command("systemctl", "stop", "v2ray").Run()
					if err != nil {
						return err
					}
					infra.GetLogger().Log("[PROC]", "stop v2ray success")
					return nil
				},
			},
		},
	}
}
