package env

import (
	"fmt"
	"log"
	"os"
)

var missingEnv = make([]string, 0)

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		errMsg := fmt.Sprintf("missing environment variable %s", key)
		missingEnv = append(missingEnv, errMsg)
	}

	return value
}

func AllEnvsOrDie() {
	if len(missingEnv) > 0 {
		for _, e := range missingEnv {
			log.Println(e)
		}
		log.Fatal("Environment variable(s) missing!")
	}
}
