package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	fiber.Use(logger.New(logger.Config{}))
	fiber.Use(cors.New())

	port := os.Getenv("API_PORT")

	fiber.Post("/api/check", handlers.HandleCheck(repos))

	slog.Info("Service listening on port: " + port)

	log.Fatal(fiber.Listen(":" + port))
}
