package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhd-sdk/haveibeenrocked/internal/repository"
)

type Repositories struct {
	PasswordRepo repository.PasswordRepository
}

func CheckTableExists(pool *pgxpool.Pool, tableName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = '%s');", tableName)
	err := pool.QueryRow(context.Background(), query).Scan(&exists)
	if err != nil {
		slog.Error("Error checking table existence:", err)
		return false
	}
	return exists
}

func CreateTable(pool *pgxpool.Pool, tableName string) error {
	query := fmt.Sprintf(`
        CREATE TABLE %s (
            id SERIAL PRIMARY KEY,
            hashed_password TEXT NOT NULL
        );
    `, tableName)

	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	fmt.Println("Table created successfully:", tableName)
	return nil
}

func Init() (*Repositories, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if !CheckTableExists(pool, "compromised_passwords") {
		err := CreateTable(pool, "compromised_passwords")
		if err != nil {
			return nil, fmt.Errorf("failed to create table and import passwords: %w", err)
		}
	}

	passwordRepo := repository.NewPasswordRepository(pool)

	return &Repositories{
		PasswordRepo: passwordRepo,
	}, nil
}
