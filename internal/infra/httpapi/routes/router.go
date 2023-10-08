package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(ctx *config.AppContext) {
	ctx.App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// app.Get("/swagger", swagger.HandlerDefault)
	loadUserRoutes(ctx.App, ctx.DB)
	loadAuthRoutes(ctx.App, ctx.DB, ctx.Env)
	ctx.App.All("*", func(c *fiber.Ctx) error {
		return httpstatus.NotFoundError("Not Found")
	})
}
