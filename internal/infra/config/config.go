package config

import (
	"github.com/fseda/cookbooked-api/pkg/env"
	"github.com/gofiber/fiber/v2/log"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Http struct {
		Port string
	}
}

func NewConfig() *Config {
	cfg := &Config{
		Database: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		} {
			Host:     env.GetEnvOrDie("DB_HOST"),
			Port:     env.GetEnvOrDie("DB_PORT"),
			User:     env.GetEnvOrDie("DB_USER"),
			Password: env.GetEnvOrDie("DB_PASSWORD"),
			Name:     env.GetEnvOrDie("DB_NAME"),
		},
		Http: struct {
			Port string
		} {
			Port: env.GetEnvOrDie("PORT"),
		},
	}

	log.Infow("Config loaded")

	return cfg
}
