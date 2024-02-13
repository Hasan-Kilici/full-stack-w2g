package handlers

import (
	"clean/database"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func RoomPage(c *fiber.Ctx) error {
	userID := c.Cookies("id")
	user, err := database.FindUser(userID)
	if err != nil {
		return c.Redirect("/")
	}

	roomID, _ := c.ParamsInt("id")
	room, err := database.FindRoom(int64(roomID))
	if err != nil {
		return c.JSON(fiber.Map{
			"message":"room not found",
		})
	}

	_ , err = database.FindRoomMember(int64(roomID), userID)
	if err != nil {
		memberData := database.RoomMember{
			RoomID		:	int64(roomID),
			UserID		:	user.ID,
			Username	:	user.Username,
			Perm		:	"member",
		}

		database.InsertRoomMember(memberData)
	}

	return c.Render("chat", fiber.Map{
		"roomName":room.Name,
		"username": user.Username,
		"userID": user.ID,
		"roomID": roomID,
	})
}