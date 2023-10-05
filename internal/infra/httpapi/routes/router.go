package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoutes(app *fiber.App, db *gorm.DB) {
	addUserRoutes(app, db)
}
