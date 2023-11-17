package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func LoadRoutes(ctx *config.AppContext) {

	loadSwagger(ctx.App)
	loadHealthCheck(ctx.App)
	
	loadUserRoutes(ctx.App, ctx.DB, ctx.Env)
	loadAuthRoutes(ctx.App, ctx.DB, ctx.Env)
	loadRecipeRoutes(ctx.App, ctx.DB, ctx.Env)
	loadUnitRoutes(ctx.App, ctx.DB, ctx.Env)
	loadIngredientRoutes(ctx.App, ctx.DB, ctx.Env)

	ctx.App.All("*", func(c *fiber.Ctx) error {
		return httpstatus.NotFoundError("Route not found")
	})

	log.Info("🛣️  Routes loaded")
}
