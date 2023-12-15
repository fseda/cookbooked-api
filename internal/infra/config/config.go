package config

import (
	"os"

	"github.com/fseda/cookbooked-api/pkg/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
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
	Github struct {
		ClientID     string
		ClientSecret string
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
			ServerPort:   env.GetEnv("PORT"),
			JWTSecretKey: []byte(env.GetEnv("JWT_SECRET_KEY")),
		},
		Github: struct {
			ClientID     string
			ClientSecret string
		}{
			ClientID:     env.GetEnv("GITHUB_CLIENT_ID"),
			ClientSecret: env.GetEnv("GITHUB_CLIENT_SECRET"),
		},
	}

	env.AllEnvsOrDie()

	log.Info("🔧 Config loaded")

	return cfg
}

func LoadDevEnvironment() {
	if isDevelopment() {
		if err := godotenv.Load(); err != nil {
			log.Warn(err)
		}
	}
}

func isDevelopment() bool {
	env := os.Getenv("GO_ENV")
	return env != "production" && env != "deploy"
}
