package infra

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Api struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type User struct {
	Default        UserDefault `yaml:"default"`
	Storage        string      `yaml:"storage"`
	NotifyTemplate string      `yaml:"notifyTemplate"`
}

type UserDefault struct {
	Tag     string `yaml:"tag"`
	Level   int    `yaml:"level"`
	AlterId int    `yaml:"alterId"`
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

func load(filepath *string) (*Config, error) {
	c := Config{}
	data, err := ioutil.ReadFile(*filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
