package infra

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dannydd88/dd-go"
	"gopkg.in/yaml.v2"
)

type Api struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type User struct {
	Default UserDefault `yaml:"default"`
	Storage string      `yaml:"storage"`
	Notify  UserNotify  `yaml:"notify"`
	Backup  UserBackup  `yaml:"backup"`
}

type UserDefault struct {
	Tag     string `yaml:"tag"`
	Level   int    `yaml:"level"`
	AlterId int    `yaml:"alterId"`
}

type UserNotify struct {
	Type     string `yaml:"type"`
	Title    string `yaml:"title"`
	Template string `yaml:"template"`
}

type UserBackup struct {
	Bucket string `yaml:"bucket"`
}

type Smtp struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Api  Api  `yaml:"api"`
	User User `yaml:"user"`
	Smtp Smtp `yaml:"smtp"`
}

func load(path *string) (*Config, error) {
	if len(dd.StringValue(path)) == 0 {
		return nil, fmt.Errorf("empty config path")
	}

	// ). find out final full config filepath
	var p string
	if filepath.IsAbs(*path) {
		p = *path
	} else {
		paths := searchPaths(path)
		for _, sp := range paths {
			if dd.FileExists(sp) {
				p = *sp
				break
			}
		}
	}

	// ). do load config
	c := Config{}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func searchPaths(path *string) []*string {
	paths := []*string{}

	// 1). current working dir
	dir, err := os.Getwd()
	if err == nil {
		paths = append(
			paths,
			dd.String(filepath.Join(dir, *path)),
		)
	}

	// 2). /etc/v2up
	paths = append(
		paths,
		dd.String(filepath.Join("/etc/v2up", *path)),
	)

	return paths
}
