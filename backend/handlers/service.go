package handlers

import(
	"fmt"
	"strconv"

	"clean/database"
	"clean/forms"
	"clean/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	snowflake 		= 	utils.NewSnowflake()
	validate 		= 	validator.New()
	badRequest		= 	"Bad request"
	roomNotFound	=	"Room Not Found"
)

func CreateRoom(c *fiber.Ctx) error {
	token := c.Cookies("token")
	user, err := database.FindUser(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Please log in")
	}

	form := new(forms.CreateRoom)
	if err = c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(badRequest)
	}

	err = validate.Struct(form)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s alanı geçerli değil", err.Field()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	id := snowflake.GenerateID();
	roomID := id;

	roomData := database.Room{
		ID			: 	roomID,
		Name		:	form.Name,
		Description	:	form.Description,
		Public		:	form.IsPublic,
	}

	modData := database.RoomMember{
		RoomID	:	roomData.ID,
		UserID	:	user.ID,
		Username:	user.Username,
		Perm	:	"admin",
	}

	settingsData := database.RoomSetting{
		RoomID			:	roomData.ID,
		StopVideo    	:	form.StopVideo,
		ChangeVideo  	:	form.ChangeVideo,
		VideoRequest 	:	form.VideoRequest,
	}

	database.InsertRoom(roomData)
	database.InsertRoomMember(modData)
	database.InsertRoomSetting(settingsData)

	return c.JSON(fiber.Map{
		"roomID": strconv.Itoa(int(roomData.ID)),
	})
}

func ListRoomMembers(c *fiber.Ctx) error {
	roomID, _ := c.ParamsInt("id")
	members := database.ListRoomMembers(int64(roomID))

	return c.JSON(fiber.Map{
		"data": members,
	})
}

func DeleteRoomMember(c *fiber.Ctx) error {
	userID := c.Params("userID")
	roomID, _ := c.ParamsInt("roomID")

	err := database.DeleteRoomMember(int64(roomID), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User Not Found")
	}

	c.SendString("User Deleted")
	return nil
}

func FindUser(c *fiber.Ctx) error {
	token := c.Params("token")

	user, err := database.FindUser(token)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User Not found")
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func FindRoom(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id");
	room, err := database.FindRoom(int64(id));
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(roomNotFound)
	}

	return c.JSON(fiber.Map{
		"data": room,
	})
}

func JoinRoom(c *fiber.Ctx) error {
	form := new(forms.JoinRoom)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(badRequest)
	}

	RoomID, err := strconv.ParseInt(form.RoomID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(badRequest)
    }

	member := database.RoomMember {
		RoomID 		:	RoomID,
		UserID 		:	form.UserID,
		Username	:	form.Username,
		Perm   		:	form.Perm,
	}

	room, err := database.FindRoom(RoomID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(roomNotFound)
	}

	database.InsertRoomMember(member)
	message := fmt.Sprintf("user joined the %s room", room.Name)

	return c.JSON(fiber.Map{
		"message": message,
	})
}

func LeaveRoom(c *fiber.Ctx) error {
	form := new(forms.LeaveRoom)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(badRequest)
	}

	RoomID, err := strconv.ParseInt(form.RoomID, 10, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(badRequest)
    }
	
	_, err = database.FindRoom(RoomID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(roomNotFound)
	}

	database.DeleteRoomMember(RoomID, form.UserID)

	return c.JSON(fiber.Map{
		"message":"message",
	})
}