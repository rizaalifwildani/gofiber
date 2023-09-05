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

type PermissionController struct {
	repository *repositories.PermissionRepository
}

func NewPermissionController(repository *repositories.PermissionRepository) *PermissionController {
	return &PermissionController{repository: repository}
}

func (c *PermissionController) CreatePermission(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.CreatePermissionRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	model := models.Permission{
		Name:        strings.ReplaceAll(strings.ToLower(req.DisplayName), " ", "-"),
		DisplayName: req.DisplayName,
	}
	err := c.repository.Create(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "permission already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "permission created successfully")
}

func (c *PermissionController) AllPermission(ctx *fiber.Ctx) error {
	filters := []repositories.FilterType{
		{Key: "name", Value: ctx.Query("name")},
		{Key: "display_name", Value: ctx.Query("displayName")},
	}
	models, _ := c.repository.FindAllPermission(filters)
	return responses.NewPermissionCollections(ctx, models)
}

func (c *PermissionController) ShowPermission(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	model, err := c.repository.FindPermission(id)
	if err != nil {
		return responses.ErrorResponse(ctx, 404, "permission not found")
	}
	return responses.NewPermissionResponse(ctx, model)
}

func (c *PermissionController) UpdatePermission(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.UpdatePermissionRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	id := ctx.Params("id")
	model := models.Permission{
		Name:        utils.GenerateSlug(req.DisplayName),
		DisplayName: req.DisplayName,
	}
	err := c.repository.Update(&model, id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "permission already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "permission updated successfully")
}
