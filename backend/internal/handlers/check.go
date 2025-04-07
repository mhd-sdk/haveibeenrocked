package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mhd-sdk/haveibeenrocked/internal/db"
)

var ctx = context.Background()

func HandleCheck(repositories *db.Repositories) fiber.Handler {
	return func(c *fiber.Ctx) error {
		prefix := c.FormValue("password")

		matchingHashess, err := repositories.PasswordRepo.FindMatching(ctx, prefix)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(matchingHashess)
	}
}
