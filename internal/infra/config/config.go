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
			Host:     env.GetEnv("PGHOST"),
			Port:     env.GetEnv("PGPORT"),
			User:     env.GetEnv("PGUSER"),
			Password: env.GetEnv("PGPASSWORD"),
			Name:     env.GetEnv("PGDATABASE"),
			Url:      env.GetEnv("DATABASE_URL"),
		},
		Http: struct {
			ServerPort   string
			JWTSecretKey []byte
		}{
			ServerPort:   env.GetEnv("SERVER_PORT"),
			JWTSecretKey: []byte(env.GetEnv("JWT_SECRET_KEY")),
		},
	}

	env.AllEnvsOrDie()

	log.Info("Config loaded")

	return cfg
}
