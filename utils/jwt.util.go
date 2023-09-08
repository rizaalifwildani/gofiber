package utils

import (
	"os"
	"strings"

	"bitbucket.org/rizaalifofficial/gofiber/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ClaimsJWT(c *fiber.Ctx) (string, *jwt.Token, *configs.JWTConfig, bool) {
	tokenString := c.Get("Authorization")
	parsedToken := strings.ReplaceAll(tokenString, "Bearer ", "")
	if tokenString == "" {
		return parsedToken, nil, nil, false
	}

	// Parse the JWT token
	jwt, err := jwt.ParseWithClaims(parsedToken, &configs.JWTConfig{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return parsedToken, nil, nil, false
	}
	claims, ok := jwt.Claims.(*configs.JWTConfig)

	return parsedToken, jwt, claims, ok
}
