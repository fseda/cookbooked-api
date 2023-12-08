package server

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/database"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/routes"
	"github.com/gofiber/fiber/v2"
)

func BuildServer(env *config.Config) (*fiber.App, func(), error) {
	db, err := database.BootstrapDB(env)
	if err != nil {
		return nil, nil, err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler:          httpstatus.ErrorHandler,
		AppName:               "Cookbooked API",
		DisableStartupMessage: true,
	})

	appContext := config.NewAppContext(app, db, env)
	routes.LoadRoutes(appContext)

	return app, func() {
		database.CloseDB(db)
	}, nil
}
