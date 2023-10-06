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
		Url      string
	}
	Http struct {
		ServerPort   string
		JWTSecretKey []byte
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
			Url      string
		}{
			Host:     env.GetEnvOrDie("PGHOST"),
			Port:     env.GetEnvOrDie("PGPORT"),
			User:     env.GetEnvOrDie("PGUSER"),
			Password: env.GetEnvOrDie("PGPASSWORD"),
			Name:     env.GetEnvOrDie("PGDATABASE"),
			Url:      env.GetEnvOrDie("DATABASE_URL"),
		},
		Http: struct {
			ServerPort   string
			JWTSecretKey []byte
		}{
			ServerPort:   env.GetEnvOrDie("PORT"),
			JWTSecretKey: []byte(env.GetEnvOrDie("JWT_SECRET_KEY")),
		},
	}

	log.Info("Config loaded")

	return cfg
}
