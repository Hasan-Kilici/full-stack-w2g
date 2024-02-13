package main

import (
	"clean/database"
	"clean/middleware"
	"clean/routes"

  	"github.com/gofiber/fiber/v2"
	"github.com/goccy/go-json"
)

func main() {
	database.CreateTables()
    app := fiber.New(fiber.Config{
		JSONEncoder : 	json.Marshal,
        JSONDecoder : 	json.Unmarshal,
    })

	app.Static("/", "./public")
	
	app.Use(middleware.Logger)
	app.Use(middleware.Compress)
	app.Use(middleware.Security)

	api := app.Group("/api", middleware.RateLimit, middleware.Jwt)
	auth := app.Group("/auth", middleware.Cors)
	server := app.Group("/server", middleware.Cors)
	service := app.Group("/service", middleware.Cors)
	ws := app.Group("/ws", middleware.Cors)

	routes.Api(api)
	routes.Auth(auth)
	routes.Server(server)
	routes.Service(service)
	routes.Websocket(ws)

	app.Use(middleware.NotFound)

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}