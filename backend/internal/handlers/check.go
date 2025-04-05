package handlers

import "github.com/gofiber/fiber/v2"

func HandleCheck(c *fiber.Ctx) error {
	helloWorld := map[string]string{
		"message": "Hello, World!",
		"status":  "ok",
	}
	return c.JSON(helloWorld)
}
