package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	c.Status(404).JSON(fiber.Map{
		"code":    404,
		"message": "404: Not Found",
	})

	return nil
}