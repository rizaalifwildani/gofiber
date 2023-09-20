package utils

import (
	"os"
	"strings"
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/configs"
	"bitbucket.org/rizaalifofficial/gofiber/static"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetToken(c *fiber.Ctx) string {
	tokenString := c.Get("Authorization")
	parsedToken := strings.ReplaceAll(tokenString, "Bearer ", "")
	return parsedToken
}

func ClaimsJWT(c *fiber.Ctx) (*configs.JWTConfig, bool) {
	tokenString := GetToken(c)
	if tokenString == "" {
		return nil, false
	}

	// Parse the JWT token
	jwt, err := jwt.ParseWithClaims(tokenString, &configs.JWTConfig{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, false
	}
	claims := jwt.Claims.(*configs.JWTConfig)

	return claims, jwt.Valid
}

func CheckJWT(c *fiber.Ctx) bool {
	token := GetToken(c)
	claims, ok := ClaimsJWT(c)

	// Check Redis
	val, err := GetRedis(static.REDIS_TOKEN)

	// Check Expired
	exp := time.Unix(claims.ExpiresAt, 0)
	currentTime := time.Now()

	if err != nil || token != string(*val) || exp.Before(currentTime) {
		return false
	}
	return ok
}
