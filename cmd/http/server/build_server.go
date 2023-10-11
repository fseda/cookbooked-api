package server

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/database"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BuildServer(env *config.Config) (*fiber.App, func(), error) {
	db, err := database.BootstrapDB(env)
	if err != nil {
		return nil, nil, err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: httpstatus.ErrorHandler,
		AppName:      "Cookbooked API",
		DisableStartupMessage: true,
	})

	loadGlobalMiddleware(app)

	appContext := config.NewAppContext(app, db, env)
	routes.LoadRoutes(appContext)

	return app, func() {
		database.CloseDB(db)
	}, nil
}

func loadGlobalMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())

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
