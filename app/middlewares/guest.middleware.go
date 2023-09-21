package middlewares

import (
	"os"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"github.com/gofiber/fiber/v2"
)

func Guest(c *fiber.Ctx) error {
	apiKey := c.Get("x-api-key")
	if apiKey != os.Getenv("X_API_KEY") {
		return responses.ErrorUnauthorized(c)
	}
	return c.Next()
}
