package routes

import (
	auth "clean/handlers"
	"github.com/gofiber/fiber/v2"
)

func Auth(app fiber.Router) {
	app.Get("/", auth.Redirect)
	app.Get("/callback", auth.Callback)
}