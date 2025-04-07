package handlers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/mhd-sdk/haveibeenrocked/internal/db"
)

var ctx = context.Background()

func HandleCheck(repositories *db.Repositories) fiber.Handler {
	return func(c *fiber.Ctx) error {
		password := c.FormValue("password")
		hash := sha1.Sum([]byte(password))
		hashPrefix := hex.EncodeToString(hash[:])[:5]

		isLeaked, err := repositories.PasswordRepo.CheckPasswordPrefix(ctx, hashPrefix)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(response)
	}
}
