package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PasswordRepository interface {
	CheckPassword(ctx context.Context, password string) (bool, error)
}

type passwordRepo struct {
	pool *pgxpool.Pool
}

func NewPasswordRepository(pool *pgxpool.Pool) PasswordRepository {
	return &passwordRepo{pool: pool}
}

func (r *passwordRepo) CheckPassword(ctx context.Context, password string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM compromised_passwords WHERE password = $1)", password).Scan(&exists)
	return exists, err
}
