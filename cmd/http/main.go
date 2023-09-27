package main

import (
	"os"

	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/database"
	"github.com/fseda/cookbooked-api/pkg/shutdown"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Exit code for graceful shutdown
	var exitCode int
	defer func() { os.Exit(exitCode) }()

	// If not able to get env, app logs Fatal error
	env := config.NewConfig()

	cleanup, err := run(env)

	defer cleanup()
	if err != nil {
		log.Errorf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()
}

func run(env *config.Config) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.Listen("0.0.0.0:" + env.Http.Port)
	}()

	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env *config.Config) (*fiber.App, func(), error) {
	db, err := database.BootstrapDB(env)
	if err != nil {
		return nil, nil, err
	}

	app := fiber.New(fiber.Config{})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// app.Get("/swagger", swagger.HandlerDefault)

	return app, func() {
		database.CloseDB(db)
	}, nil
}
