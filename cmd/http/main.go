package main

import (
	"os"

	"github.com/fseda/cookbooked-api/cmd/http/server"
	_ "github.com/fseda/cookbooked-api/docs"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/pkg/shutdown"

	"github.com/gofiber/fiber/v2/log"
)

//	@title						CookBooked API
//	@version					1.0
//	@description				API for CookBooked, a recipe management app.
//	@contact.name				Felipe Seda
//	@license.name				MIT
//	@BasePath					/
//
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
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
