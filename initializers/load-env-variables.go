package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: %w", err)
		return err
	}
	return nil
}
