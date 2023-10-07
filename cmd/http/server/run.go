package server

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"

	"github.com/gofiber/fiber/v2/log"
)

func Run(env *config.Config) (func(), error) {
	app, cleanup, err := BuildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := app.Listen(":"+env.Http.ServerPort); err != nil {
			log.Fatal(err)
		}
	}()

	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}
