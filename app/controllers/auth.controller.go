package controllers

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/requests"
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

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.LoginRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	res, err := c.repository.Login(req.Username, req.Password)
	if err != nil {
		return responses.ErrorValidationResponse(ctx, "invalid username or password")
	}

	return responses.NewAuthResponse(ctx, *res)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	err := c.repository.Logout(ctx)
	if err != nil {
		return responses.ErrorUnauthorized(ctx)
	}

	return responses.SuccessResponse(ctx, "logged out successfully")
}
