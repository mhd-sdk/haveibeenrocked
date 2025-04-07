package repositories

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PasswordRepository interface {
	CheckPassword(ctx context.Context, password string) (bool, error)
	CheckPasswordPrefix(ctx context.Context, prefix string) (bool, error)
}

type passwordRepo struct {
	pool *pgxpool.Pool
}

func NewPasswordRepository(pool *pgxpool.Pool) PasswordRepository {
	return &passwordRepo{pool: pool}
}

func (r *passwordRepo) CheckPassword(ctx context.Context, password string) (bool, error) {
	var storedHash, salt string
	err := r.pool.QueryRow(ctx, "SELECT hashed_password, salt FROM compromised_passwords WHERE hashed_password = $1", password).Scan(&storedHash, &salt)
	if err != nil {
		return false, err
	}

	hashedInput := sha1.Sum([]byte(password + salt))
	return hex.EncodeToString(hashedInput[:]) == storedHash, nil
}

func (r *passwordRepo) CheckPasswordPrefix(ctx context.Context, prefix string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM compromised_passwords WHERE hashed_password LIKE $1"
	err := r.pool.QueryRow(ctx, query, prefix+"%").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
