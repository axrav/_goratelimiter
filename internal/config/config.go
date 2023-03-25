package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loading the port with go dotenv

func Load() {
	// loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Get(key string) string {
	// getting the port from the .env file
	return os.Getenv(key)
}
