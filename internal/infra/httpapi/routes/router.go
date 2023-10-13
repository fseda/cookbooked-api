package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func LoadRoutes(ctx *config.AppContext) {
	// app.Get("/swagger", swagger.HandlerDefault)

	loadHealthCheck(ctx.App)
	loadUserRoutes(ctx.App, ctx.DB, ctx.Env)
	loadAuthRoutes(ctx.App, ctx.DB, ctx.Env)
	loadRecipeRoutes(ctx.App, ctx.DB, ctx.Env)

	ctx.App.All("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Info("🛣️  Routes loaded")
}
