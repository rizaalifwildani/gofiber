package middlewares

import (
	"strings"
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/static"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"github.com/gofiber/fiber/v2"
)

func InitMiddleware(c *fiber.Ctx) error {
	token, jwt, claims, ok := utils.ClaimsJWT(c)

	// Check Redis
	val, err := utils.GetRedis(static.REDIS_TOKEN)

	// Check Expired
	exp := time.Unix(claims.ExpiresAt, 0)
	currentTime := time.Now()

	if err != nil || token != string(*val) || exp.Before(currentTime) {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "invalid token")
	}

	if ok && jwt.Valid {
		// Check the requested path
		requestedPath := c.Path()
		if strings.Contains(requestedPath, "users") {
			// Check the user's role from the claims
			roles := claims.Roles
			for _, role := range roles {
				if !strings.Contains(role.Name, "super") {
					return responses.ErrorForbidden(c)
				}
			}
		}
		return c.Next()
	}

	return responses.ErrorUnauthorized(c)
}
