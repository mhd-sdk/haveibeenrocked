package db

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhd-sdk/haveibeenrocked/internal/repositories"
)

type Repositories struct {
	PasswordRepo repositories.PasswordRepository
}

func Init() (*Repositories, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	passwordRepo := repositories.NewPasswordRepository(pool, redisClient)

	return &Repositories{
		PasswordRepo: passwordRepo,
	}, nil
}
