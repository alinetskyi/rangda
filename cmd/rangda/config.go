package main

import (
	"github.com/kovetskiy/ko"
	"gopkg.in/yaml.v2"
)

type ConfigSession struct {
	SecretKeyBase []byte `required:"true"`
}

type Config struct {
	Address string `default:":8080" env:"RANGDA_ADDRESS"`
	Session ConfigSession
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	err := ko.Load(path, &config, yaml.Unmarshal)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
