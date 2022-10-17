package infra

import (
	"github.com/dannydd88/dd-go"
	"github.com/urfave/cli/v2"
)

type V2upContext struct {
	Logging dd.Logger
	Config  Config
	Mailer  *Mailer
}

var globalContext V2upContext

func Init(c *cli.Context) error {
	// ). init context
	globalContext = V2upContext{}

	// ). init logging
	globalContext.Logging = dd.NewDefaultLogger()

	// ). init config
	{
		// ). Load config
		config, err := load(dd.String(c.String("config")))
		if err != nil {
			return err
		}

		globalContext.Config = *config
	}

	// ). init mailer
	{
		globalContext.Mailer = &Mailer{
			smtpAdress:   globalContext.Config.Smtp.Address,
			smtpPort:     globalContext.Config.Smtp.Port,
			smtpUsername: globalContext.Config.Smtp.Username,
			smtpPassword: globalContext.Config.Smtp.Password,
		}
	}

	return nil
}

func GetLogger() dd.Logger {
	return globalContext.Logging
}

func GetConfig() *Config {
	return &globalContext.Config
}

func GetMailer() *Mailer {
	return globalContext.Mailer
}
