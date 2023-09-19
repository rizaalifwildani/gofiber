package configs

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTConfig struct {
	jwt.StandardClaims
	ID    uuid.UUID                `json:"id"`
	Roles []responses.RoleResponse `json:"roles"`
}
