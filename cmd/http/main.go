package main

import (
	"os"
	"strings"

	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/database"
	"github.com/fseda/cookbooked-api/pkg/shutdown"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Exit code for graceful shutdown
	var exitCode int
	defer func() { os.Exit(exitCode) }()
	
	// If not able to get env, app logs Fatal error
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("Could not load environment variables")
	}
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodHead,
		}, ","),
	}))
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

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// app.Get("/swagger", swagger.HandlerDefault)

	return app, func() {
		database.CloseDB(db)
	}, nil
}
