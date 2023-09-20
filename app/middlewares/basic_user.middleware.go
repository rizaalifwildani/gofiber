package middlewares

import (
	"os"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func BasicUser() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return responses.ErrorUnauthorized(ctx)
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			_, ok := utils.ClaimsJWT(c)

			if ok {
				return c.Next()
			}

			return responses.ErrorResponse(c, fiber.StatusUnauthorized, "invalid token")
		},
	})
}
