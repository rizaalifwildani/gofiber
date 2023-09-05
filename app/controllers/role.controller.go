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

type RoleController struct {
	repository *repositories.RoleRepository
}

func NewRoleController(repository *repositories.RoleRepository) *RoleController {
	return &RoleController{repository: repository}
}

func (c *RoleController) CreateRole(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.CreateRoleRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	model := models.Role{
		Name:        strings.ReplaceAll(strings.ToLower(req.DisplayName), " ", "-"),
		DisplayName: req.DisplayName,
		Permissions: req.Permissions,
	}
	err := c.repository.Create(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "role already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "role created successfully")
}

func (c *RoleController) AllRole(ctx *fiber.Ctx) error {
	filters := []repositories.FilterType{
		{Key: "name", Value: ctx.Query("name")},
		{Key: "display_name", Value: ctx.Query("displayName")},
	}
	models, _ := c.repository.FindAllRole(filters)
	return responses.NewRoleCollections(ctx, models)
}

func (c *RoleController) ShowRole(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	model, err := c.repository.FindRole(id)
	if err != nil {
		return responses.ErrorResponse(ctx, 404, "role not found")
	}
	return responses.NewRoleResponse(ctx, model)
}

func (c *RoleController) UpdateRole(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.UpdateRoleRequest{}
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
	model := models.Role{
		ID:          parsedUUID,
		Name:        utils.GenerateSlug(req.DisplayName),
		DisplayName: req.DisplayName,
		Permissions: req.Permissions,
	}
	err := c.repository.UpdateRole(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "role already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "role updated successfully")
}
