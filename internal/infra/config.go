package infra

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type V2rayConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	V2ray V2rayConfig `yaml:"v2ray"`
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
