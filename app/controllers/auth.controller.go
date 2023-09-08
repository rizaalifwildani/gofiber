package controllers

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	repository *repositories.AuthRepository
}

func NewAuthController(repository *repositories.AuthRepository) *AuthController {
	return &AuthController{repository: repository}
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	err := c.repository.Logout(ctx)
	if err != nil {
		return responses.ErrorUnauthorized(ctx)
	}

	return responses.SuccessResponse(ctx, "logged out successfully")
}
