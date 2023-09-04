package responses

import (
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
)

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}

func NewAuthResponse(ctx *fiber.Ctx, m models.UserAuth) error {
	data := AuthResponse{
		Token:     m.Token,
		ExpiredAt: m.ExpiredAt.UnixNano() / int64(time.Millisecond),
	}
	return SuccessResponse(ctx, data)
}
