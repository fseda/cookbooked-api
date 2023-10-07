package main

import (
	"os"

	"github.com/fseda/cookbooked-api/cmd/http/server"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/pkg/shutdown"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// Exit code for graceful shutdown
	var exitCode int
	defer func() { os.Exit(exitCode) }()

	config.LoadDevEnvironment()

	// If not able to get env, app logs Fatal error
	env := config.NewConfig()

	cleanup, err := server.Run(env)

	defer cleanup()
	if err != nil {
		log.Errorf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()
}
