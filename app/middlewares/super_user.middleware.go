package middlewares

import (
	"os"
	"strings"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SuperUser() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return responses.ErrorUnauthorized(ctx)
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			claims, ok := utils.ClaimsJWT(c)

			if ok {
				// Check the user's role from the claims
				roles := claims.User.Roles
				for _, role := range roles {
					if strings.Contains(role.Role.Name, "root") || strings.Contains(role.Role.Name, "super") || strings.Contains(role.Role.Name, "pusat") {
						return c.Next()
					}
				}
				return responses.ErrorForbidden(c)
			}

			return responses.ErrorResponse(c, fiber.StatusUnauthorized, "invalid token")
		},
	})
}
