package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// https://www.codingexplorations.com/blog/mastering-the-repository-pattern-in-go-a-comprehensive-guide
type PasswordRepository interface {
	FindMatching(ctx context.Context, prefix string) ([]string, error)
}

type passwordRepo struct {
	pool *pgxpool.Pool
}

func NewPasswordRepository(pool *pgxpool.Pool) PasswordRepository {
	return &passwordRepo{pool: pool}
}

func (r *passwordRepo) FindMatching(ctx context.Context, prefix string) ([]string, error) {
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

	return passwords, nil
}
