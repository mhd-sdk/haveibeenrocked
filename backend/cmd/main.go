package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	fiber.Use(logger.New(logger.Config{}))

	fiber.Post("/api/v1", handlers.HandleCheck)
	log.Fatal(fiber.Listen(":" + os.Getenv("API_PORT")))
}
