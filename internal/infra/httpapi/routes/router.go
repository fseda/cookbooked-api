package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	middlewares "github.com/fseda/cookbooked-api/internal/infra/httpapi/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoadRoutes(ctx *config.AppContext) {
	// Middlewares
	loadGlobalMiddleware(ctx.App)


	// Routes
	loadSwagger(ctx.App)
	loadHealthCheck(ctx.App)
	
	loadUserRoutes(ctx.App, ctx.DB, ctx.Env)
	loadAuthRoutes(ctx.App, ctx.DB, ctx.Env)
	loadRecipeRoutes(ctx.App, ctx.DB, ctx.Env)

	ctx.App.All("*", func(c *fiber.Ctx) error {
		return httpstatus.NotFoundError("Route not found")
	})

	log.Info("🛣️  Routes loaded")
}

func loadGlobalMiddleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())
	app.Use(middlewares.RateLimiter())

	app.Hooks().OnListen(func(listenData fiber.ListenData) error {
		if fiber.IsChild() {
			return nil
		}
		scheme := "http"
		if listenData.TLS {
			scheme = "https"
		}

		log.Infof("🚀 Server listening on %v://%v:%v", scheme, listenData.Host, listenData.Port)
		return nil
	})

	log.Info("🛡️  Global middleware loaded")
}
