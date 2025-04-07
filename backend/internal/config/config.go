package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrMissingEnvVars = errors.New("missing environment variables")
)

// Load env variables from .env file if it exists, otherwise use system environment variables
func LoadEnv() (err error) {
	// this will NEVER overrride an env var that already exist
	godotenv.Load()

	apiPort := os.Getenv("API_PORT")
	redisHost := os.Getenv("REDIS_HOST")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if apiPort == "" || redisHost == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {

		return ErrMissingEnvVars
	}

	os.Setenv("DATABASE_URL", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName))

	return nil
}
