package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// LoadPasswordsFromFile charge les mots de passe depuis un fichier texte.
func LoadPasswordsFromFile(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	passwords := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			passwords = append(passwords, line)
		}
	}
	return passwords, nil
}

// HashPassword hache un mot de passe avec bcrypt.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("erreur lors du hachage du mot de passe: %w", err)
	}
	return string(hashedBytes), nil
}

// InsertHashedPasswords insère les mots de passe hachés dans la base de données.
func InsertHashedPasswords(pool *pgxpool.Pool, tableName string, passwords []string) error {
	query := fmt.Sprintf("INSERT INTO %s (hashed_password) VALUES ($1);", tableName)

	for _, password := range passwords {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			return err
		}

		_, err = pool.Exec(context.Background(), query, hashedPassword)
		if err != nil {
			return fmt.Errorf("erreur lors de l'insertion du mot de passe: %w", err)
		}
	}

	fmt.Println("Mots de passe importés avec succès.")
	return nil
}
