package configs

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/golang-jwt/jwt"
)

type JWTConfig struct {
	jwt.StandardClaims
	User models.User `json:"user"`
}
