package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

// https://www.codingexplorations.com/blog/mastering-the-repository-pattern-in-go-a-comprehensive-guide
type PasswordRepository interface {
	FindMatching(ctx context.Context, prefix string) ([]string, error)
}

type passwordRepo struct {
	pool  *pgxpool.Pool
	redis *redis.Client
}

func NewPasswordRepository(pool *pgxpool.Pool, redisClient *redis.Client) PasswordRepository {
	return &passwordRepo{
		pool:  pool,
		redis: redisClient,
	}
}

func (r *passwordRepo) FindMatching(ctx context.Context, prefix string) ([]string, error) {
	cacheKey := prefix
	cachedResult, err := r.redis.Get(ctx, cacheKey).Result()

	// If found in cache and no error, deserialize and return
	if err == nil {
		var passwords []string
		if err := json.Unmarshal([]byte(cachedResult), &passwords); err == nil {
			return passwords, nil
		}
	}

	var passwords []string
	query := "SELECT hashed_password FROM compromised_passwords WHERE hashed_password LIKE $1"
	rows, err := r.pool.Query(ctx, query, prefix+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var password string
		if err := rows.Scan(&password); err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Cache the result for future requests
	if len(passwords) > 0 {
		jsonData, err := json.Marshal(passwords)
		if err == nil {
			r.redis.Set(ctx, cacheKey, jsonData, time.Hour)
		}
	}

	return passwords, nil
}
