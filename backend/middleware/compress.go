package middleware

import(
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2"
)

func Compress(c *fiber.Ctx) error {
	compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
	return c.Next()
}