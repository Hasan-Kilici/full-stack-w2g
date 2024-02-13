package middleware

import "github.com/gofiber/fiber/v2/middleware/cors"

var Cors = cors.New(cors.Config{
        AllowOrigins		:     	"http://localhost:5173",
        AllowHeaders		:     	"Origin, Content-Type, Accept",
        AllowCredentials	: 	true,
})