package handlers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/mhd-sdk/haveibeenrocked/internal/anssi"
	"github.com/mhd-sdk/haveibeenrocked/internal/db"
)

type PasswordCheckResponse struct {
	IsLeaked        bool                   `json:"isLeaked"`
	Recommendations []anssi.Recommendation `json:"recommendations"`
}

var ctx = context.Background()

func HandleCheck(repositories *db.Repositories) fiber.Handler {
	return func(c *fiber.Ctx) error {
		password := c.FormValue("password")
		hash := sha1.Sum([]byte(password))
		hashPrefix := hex.EncodeToString(hash[:])[:5]

		checkResult := anssi.CheckPassword(password)

		isLeaked, err := repositories.PasswordRepo.CheckPasswordPrefix(ctx, hashPrefix)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		response := PasswordCheckResponse{
			IsLeaked:        isLeaked,
			Recommendations: checkResult.Missing,
		}

		return c.JSON(response)
	}
}
