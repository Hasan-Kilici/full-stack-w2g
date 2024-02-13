package routes

import (
	"clean/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Websocket(app fiber.Router) {
	app.Get("/:room", websocket.New(func(c *websocket.Conn) {
		room := c.Params("room")
		ws.HandleRooms(c, room)
	}))
}