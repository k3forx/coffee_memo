package config

import (
	"github.com/kelseyhightower/envconfig"
)

var cfg Config

type Config struct {
	MainAppPortNumber int `envconfig:"MAIN_APP_PORT_NUMBER"`
}

func LoadConfig() error {
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}
	return nil
}

func GetMainAppPortNumber() int {
	return cfg.MainAppPortNumber
}
