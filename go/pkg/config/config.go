package config

import "github.com/spf13/viper"

type Config struct {
	Test string `mapstructure:"TEST"`
}

func LoadConfig() error {
	var cfg Config
	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	return err
	// }

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}
