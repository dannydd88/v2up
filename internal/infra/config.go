package infra

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dannydd88/dd-go"
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

func ensurePath(path *string) (*string, error) {
	if len(dd.Val(path)) == 0 {
		return nil, fmt.Errorf("empty config path")
	}

	// ). check if it is an absolute path
	if filepath.IsAbs(*path) {
		return path, nil
	}

	// ). build searching path
	paths := []*string{}
	{
		// 1). current working dir
		dir, err := os.Getwd()
		if err == nil {
			paths = append(
				paths,
				dd.Ptr(filepath.Join(dir, dd.Val(path))),
			)
		}

		// 2). /etc/v2up
		paths = append(
			paths,
			dd.Ptr(filepath.Join("/etc/v2up", dd.Val(path))),
		)
	}

	// ). foreach paths, return path if it is existed
	for _, sp := range paths {
		if dd.FileExists(sp) {
			return sp, nil
		}
	}

	// ). return error, because cannot find path
	return nil, fmt.Errorf("cannot find file -> %s", dd.Val(path))
}
