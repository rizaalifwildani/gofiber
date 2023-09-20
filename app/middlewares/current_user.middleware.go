package middlewares

import (
	"os"
	"strings"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func CurrentUser() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return responses.ErrorUnauthorized(ctx)
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			claims, ok := utils.ClaimsJWT(c)

			if ok {
				// Check the user's role from the claims
				for _, v := range claims.User.Roles {
					if strings.Contains(v.Name, "root") {
						return c.Next()
					}
				}
				if c.Params("id") == claims.User.ID.String() {
					return c.Next()
				}
				return responses.ErrorForbidden(c)
			}

			return responses.ErrorResponse(c, fiber.StatusUnauthorized, "invalid token")
		},
	})
}
