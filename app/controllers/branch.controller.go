package controllers

import (
	"strings"

	"bitbucket.org/rizaalifofficial/gofiber/app/requests"
	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
)

type BranchController struct {
	repository *repositories.BranchRepository
}

func NewBranchController(repository *repositories.BranchRepository) *BranchController {
	return &BranchController{repository: repository}
}

func (c *BranchController) CreateBranch(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.CreateBranchRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	model := models.Branch{
		Name:        req.Name,
		Code:        req.Code,
		Address:     req.Address,
		Description: req.Description,
	}
	err := c.repository.Create(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "branch already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "branch created successfully")
}

func (c *BranchController) AllBranch(ctx *fiber.Ctx) error {
	filters := []repositories.FilterType{
		{Key: "name", Value: ctx.Query("name")},
		{Key: "code", Value: ctx.Query("code")},
		{Key: "address", Value: ctx.Query("address")},
	}
	models, _ := c.repository.FindAllBranch(filters)
	return responses.NewBranchCollections(ctx, models)
}

func (c *BranchController) ShowBranch(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	model, err := c.repository.FindBranch(id)
	if err != nil {
		return responses.ErrorResponse(ctx, 404, "branch not found")
	}
	return responses.NewBranchResponse(ctx, model)
}

func (c *BranchController) UpdateBranch(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.UpdateBranchRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	id := ctx.Params("id")
	model := models.Branch{
		Name:        req.Name,
		Code:        req.Code,
		Address:     req.Address,
		Description: req.Description,
	}
	err := c.repository.Update(&model, id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "branch already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "branch updated successfully")
}
