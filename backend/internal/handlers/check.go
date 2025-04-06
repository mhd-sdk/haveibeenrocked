package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/mhd-sdk/haveibeenrocked/internal/anssi"
)

type PasswordCheckResponse struct {
	IsLeaked        bool                   `json:"isLeaked"`
	Recommendations []anssi.Recommendation `json:"recommendations"`
}

var ctx = context.Background()

func HandleCheck(db *pgx.Conn, redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		password := c.FormValue("password")

		// Vérifier le cache Redis
		cacheKey := "password:" + password
		cachedResult, err := redisClient.Get(ctx, cacheKey).Result()
		if err == nil {
			// Si trouvé dans le cache, renvoyer la réponse
			return c.JSON(cachedResult)
		}

		// Vérifier les recommandations ANSSI
		checkResult := anssi.CheckPassword(password)

		// Vérifier dans la base de données si le mot de passe est compromis
		var isLeaked bool
		err = db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM compromised_passwords WHERE password = $1)", password).Scan(&isLeaked)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}

		// Créer la réponse
		response := PasswordCheckResponse{
			IsLeaked:        isLeaked,
			Recommendations: checkResult.Missing,
		}

		// Mettre en cache le résultat dans Redis
		redisClient.Set(ctx, cacheKey, response, 0)

		return c.JSON(response)
	}
}
