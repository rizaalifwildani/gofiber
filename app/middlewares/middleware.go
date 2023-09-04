package middlewares

import (
	"log"
	"os"
	"strings"
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func InitMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return responses.ErrorUnauthorized(c)
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(strings.ReplaceAll(tokenString, "Bearer ", ""), &configs.JWTConfig{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return responses.ErrorUnauthorized(c)
	}
	claims, ok := token.Claims.(*configs.JWTConfig)

	log.Println(claims.ExpiresAt)

	exp := time.Unix(claims.ExpiresAt, 0)
	currentTime := time.Now()

	// Compare the expiration time to the current time
	if exp.Before(currentTime) {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "token has expired")
	}

	if ok && token.Valid {
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
