package controllers

import (
	"strings"

	"bitbucket.org/rizaalifofficial/gofiber/app/requests"
	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	repository *repositories.UserRepository
}

func NewUserController(repository *repositories.UserRepository) *UserController {
	return &UserController{repository: repository}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.CreateUserRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	model := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Roles:     req.Roles,
	}
	authModel := models.UserAuth{
		Password: req.Password,
	}
	err := c.repository.CreateUser(&model, &authModel)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "user already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "user created successfully")
}

func (c *UserController) AllUser(ctx *fiber.Ctx) error {
	filters := []repositories.FilterType{
		{Key: "username", Value: ctx.Query("username")},
		{Key: "email", Value: ctx.Query("email")},
		{Key: "phone", Value: ctx.Query("phone")},
		{Key: "first_name", Value: ctx.Query("firstName")},
		{Key: "last_name", Value: ctx.Query("lastName")},
	}
	models, _ := c.repository.FindAllUser(filters)
	return responses.NewUserCollections(ctx, models)
}

func (c *UserController) ShowUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	model, err := c.repository.FindUser(id)
	if err != nil {
		return responses.ErrorResponse(ctx, 404, "user not found")
	}
	return responses.NewUserResponse(ctx, model)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.UpdateUserRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	id := ctx.Params("id")
	parsedUUID := utils.GenerateUUID(id)
	model := models.User{
		ID:        parsedUUID,
		Email:     req.Email,
		Phone:     req.Phone,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Roles:     req.Roles,
	}
	authModel := models.UserAuth{
		UserID:    parsedUUID,
		ExpiredAt: nil,
		Token:     "",
		Password:  req.Password,
	}
	err := c.repository.UpdateUser(&model, &authModel)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "user already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "user updated successfully")
}
