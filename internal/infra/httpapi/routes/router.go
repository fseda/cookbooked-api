package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(ctx *config.AppContext) {
	// app.Get("/swagger", swagger.HandlerDefault)
	
	loadHealthCheck(ctx.App)
	loadUserRoutes(ctx.App, ctx.DB, ctx.Env)
	loadAuthRoutes(ctx.App, ctx.DB, ctx.Env)
	
	ctx.App.All("*", func(c *fiber.Ctx) error {
		return httpstatus.NotFoundError("Not Found")
	})
}
