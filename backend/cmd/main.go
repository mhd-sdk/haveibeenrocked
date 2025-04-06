package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/lmittmann/tint"
	"github.com/mhd-sdk/haveibeenrocked/internal/config"
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

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiber.Use(logger.New(logger.Config{}))

	fiber.Post("/api/check", handlers.HandleCheck(db, redisClient))

	slog.Info("Service listening on port: " + os.Getenv("API_PORT"))

	log.Fatal(fiber.Listen(":" + os.Getenv("API_PORT")))
}
