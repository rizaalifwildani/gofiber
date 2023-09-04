package configs

import (
	"os"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTConfig struct {
	jwt.StandardClaims
	ID    uuid.UUID                `json:"id"`
	Roles []responses.RoleResponse `json:"roles"`
}

func InitJWT() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return responses.ErrorUnauthorized(ctx)
		},
	})
}
