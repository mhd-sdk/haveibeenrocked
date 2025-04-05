package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrMissingEnvVars = errors.New("missing environment variables")
)

// Load env variables from .env file if it exists, otherwise use system environment variables
func LoadEnv() (err error) {
	godotenv.Load()

	apiPort := os.Getenv("API_PORT")
	redisHost := os.Getenv("REDIS_HOST")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if apiPort == "" || redisHost == "" || dbHost == "" || dbPort == "" {
		return ErrMissingEnvVars
	}

	return nil
}
