package routes

import (
	service "clean/handlers"
	"github.com/gofiber/fiber/v2"
)

func Service(app fiber.Router) {
	app.Post("/create/room", service.CreateRoom)
	app.Post("/join/room", service.JoinRoom)
	app.Post("/leave/room", service.LeaveRoom)
	
	app.Delete("/delete/:roomID/:userID",service.DeleteRoomMember)

	app.Get("/members/:id", service.ListRoomMembers)
	app.Get("/find/user/:token", service.FindUser)
	app.Get("/find/room/:id", service.FindRoom)
}