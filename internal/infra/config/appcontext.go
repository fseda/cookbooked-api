package config

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppContext struct {
	App *fiber.App
	DB  *gorm.DB
	Env *Config
}

func NewAppContext(app *fiber.App, db *gorm.DB, env *Config) *AppContext {
	return &AppContext{
		app,
		db,
		env,
	}
}
