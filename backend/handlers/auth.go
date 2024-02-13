package handlers

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"clean/utils"
	"clean/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"github.com/goccy/go-json"
)

var (
	state = "random"
	conf *oauth2.Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf = &oauth2.Config{
		RedirectURL		:  "http://localhost:3000/auth/callback",
		ClientID		:   os.Getenv("client_id"),
		ClientSecret	: 	os.Getenv("client_secret"),
		Scopes			:   []string{discord.ScopeIdentify},
		Endpoint		:   discord.Endpoint,
	}
}

func Redirect(c *fiber.Ctx) error {
	return c.Redirect(conf.AuthCodeURL(state), fiber.StatusTemporaryRedirect)
}

func Callback(c *fiber.Ctx) error {
	if c.Query("state") != state {
		return c.Status(fiber.StatusBadRequest).SendString("State does not match.")
	}

	token, err := conf.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	res, err := conf.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil || res.StatusCode != 200 {
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString(res.Status)
		}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var user database.User

	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"

	member , err := database.FindUserById(user.ID)
	if err != nil {
		user.Token = utils.GenerateToken(user.ID)
		cookie.Value = user.Token

		database.InsertUser(user)
	
		c.Cookie(cookie)
		return c.Redirect("http://localhost:5173")
	}

	cookie.Value = member.Token
	c.Cookie(cookie)

	return c.Redirect("http://localhost:5173")
}