package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	quit := make(chan os.Signal)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
