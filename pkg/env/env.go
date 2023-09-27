package env

import (
	"fmt"
	"log"
	"os"
)

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)
	
	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		log.Fatal(err)
	}

	return value
}
