package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitCors() func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins:     "*", // Set to your allowed origins, "*" allows all origins
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,     // Set to true if you want to allow credentials (cookies, HTTP authentication) cross-origin
		MaxAge:           3600 * 24, // Cache preflight request for 1 hour (optional)
	})
}
