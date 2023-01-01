package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


// MustGet will return the env or throw an error if not present.
func MustGet(key string) string {
	value := os.Getenv(key)

	if value == "" && key != "PORT" {
		log.Printf("Env key missing" + key)
	}

	return value
}

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
}
