package infra

import (
	"github.com/dannydd88/dd-go"
	"github.com/urfave/cli/v2"
)

type V2upContext struct {
	Logging dd.LevelLogger
	Config  *Config
	Mailer  *Mailer
}

var globalContext V2upContext

func AppInit(c *cli.Context) error {
	// ). init context
	globalContext = V2upContext{}

	// ). init logging
	globalContext.Logging = dd.NewLevelLogger(dd.INFO)

	return nil
}

func CommandInit(c *cli.Context) error {
	// ). init config
	{
		globalContext.Config = &Config{}
		// ). Load config
		path, err := ensurePath(dd.Ptr(c.String("config")))
		if err != nil {
			return err
		}
		err = dd.NewYAMLLoader[Config](path).Load(globalContext.Config)
		if err != nil {
			return err
		}
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

func GetLogger() dd.LevelLogger {
	return globalContext.Logging
}

func GetConfig() *Config {
	return globalContext.Config
}

func GetMailer() *Mailer {
	return globalContext.Mailer
}
