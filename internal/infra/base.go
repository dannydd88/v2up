package infra

import (
	"os"
	"path/filepath"

	"github.com/dannydd88/gobase/pkg/base"
	"github.com/urfave/cli/v2"
)

type V2upContext struct {
	Logging base.Logger
	Config  Config
	Mailer  *Mailer
}

var globalContext V2upContext

func Init(c *cli.Context) error {
	// ). init context
	globalContext = V2upContext{}

	// ). init logging
	globalContext.Logging = base.NewDefaultLogger()

	// ). init config
	var config *Config
	{
		// ). Get current dir
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		// ). Load config
		configPath := c.String("config")
		if !filepath.IsAbs(configPath) {
			configPath = filepath.Join(dir, configPath)
		}
		config, err = load(base.String(configPath))
		if err != nil {
			return err
		}

		globalContext.Config = *config
	}

	// ). init mailer
	{
		globalContext.Mailer = &Mailer{
			smtpAdress:   config.Smtp.Address,
			smtpPort:     config.Smtp.Port,
			smtpUsername: config.Smtp.Username,
			smtpPassword: config.Smtp.Password,
		}
	}

	return nil
}

func GetLogger() base.Logger {
	return globalContext.Logging
}

func GetConfig() *Config {
	return &globalContext.Config
}

func GetMailer() *Mailer {
	return globalContext.Mailer
}
