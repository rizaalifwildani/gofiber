package controllers

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/requests"
	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
)

type GuestController struct {
	repository *repositories.GuestRepository
}

func NewGuestController(repository *repositories.GuestRepository) *GuestController {
	return &GuestController{repository: repository}
}

func (c *GuestController) Login(ctx *fiber.Ctx) error {
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
