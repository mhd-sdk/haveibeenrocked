package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lmittmann/tint"
	"github.com/mhd-sdk/haveibeenrocked/internal/config"
	"github.com/mhd-sdk/haveibeenrocked/internal/db"
	"github.com/mhd-sdk/haveibeenrocked/internal/handlers"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	slog.Info("Starting backend service...")

	err := config.LoadEnv()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	repos, err := db.Init()
	if err != nil {
		log.Fatal("Failed to initialize database and repositories:", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiber.Use(logger.New(logger.Config{}))

	fiber.Post("/api/check", handlers.HandleCheck(repos, redisClient))

	slog.Info("Service listening on port: " + os.Getenv("API_PORT"))

	log.Fatal(fiber.Listen(":" + os.Getenv("API_PORT")))
}
