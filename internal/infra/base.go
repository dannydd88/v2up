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
}

var globalContext V2upContext

func Init(c *cli.Context) error {
	// ). init context
	globalContext = V2upContext{}

	// ). init logging
	globalContext.Logging = base.NewDefaultLogger()

	// ). init config
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
		config, err := load(base.String(configPath))
		if err != nil {
			return err
		}

		globalContext.Config = *config
	}

	return nil
}

func GetLogger() base.Logger {
	return globalContext.Logging
}

func GetConfig() *Config {
	return &globalContext.Config
}
