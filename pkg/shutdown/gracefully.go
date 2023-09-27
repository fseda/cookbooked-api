package shutdown

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
)

func Gracefully() {
	quit := make(chan os.Signal)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Infof("\nGracefully shutdown (%v)\n", <-quit)
}
