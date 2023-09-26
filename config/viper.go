package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT          string `mapstructure:"PORT"`
	POSTGRES_URI  string `mapstructure:"POSTGRES_URI"`
	POSTGRES_NAME string `mapstructure:"POSTGRES_NAME"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			PORT:          os.Getenv("PORT"),
			POSTGRES_URI:  os.Getenv("POSTGRES_URI"),
			POSTGRES_NAME: os.Getenv("POSTGRES_NAME"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// Validate config
	if config.POSTGRES_URI == "" {
		err = errors.New("POSTGRES_URI is required")
		return
	}

	if config.POSTGRES_NAME == "" {
		err = errors.New("POSTGRES_NAME is required")
		return
	}

	return
}
