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

type MemberController struct {
	repository *repositories.MemberRepository
}

func NewMemberController(repository *repositories.MemberRepository) *MemberController {
	return &MemberController{repository: repository}
}

func (c *MemberController) CreateMember(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.CreateMemberRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return responses.ErrorValidationResponse(ctx, err.Error())
	}
	errors := requests.NewValidatorRequest(ctx, &req)
	if len(errors) > 0 {
		return responses.ErrorValidationResponse(ctx, errors)
	}

	/* === PROCESS === */
	birthDate, birthDateErr := utils.DateUtilYYYYMMDD(req.Birthdate)
	if birthDateErr != nil {
		return responses.ErrorValidationResponse(ctx, "invalid birthdate")
	}
	model := models.Member{
		Phone:          req.Phone,
		Email:          req.Email,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		IdentityNumber: req.IdentityNumber,
		PlaceOfBirth:   req.PlaceOfBirth,
		Birthdate:      *birthDate,
		Gender:         req.Gender,
		Nationality:    req.Nationality,
		Address:        req.Address,
		PostalCode:     req.PostalCode,
		HomePhone:      req.HomePhone,
		OfficePhone:    req.OfficePhone,
		Education:      req.Education,
		Branches:       req.Branches,
	}
	err := c.repository.Create(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "member already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "member created successfully")
}

func (c *MemberController) AllMember(ctx *fiber.Ctx) error {
	filters := []repositories.FilterType{
		{Key: "phone", Value: ctx.Query("phone")},
		{Key: "email", Value: ctx.Query("email")},
		{Key: "firstName", Value: ctx.Query("firstName")},
		{Key: "lastName", Value: ctx.Query("lastName")},
		{Key: "branch", Value: ctx.Query("branch")},
	}
	models, _ := c.repository.FindAllMember(filters)
	return responses.NewMemberCollections(ctx, models)
}

func (c *MemberController) ShowMember(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	model, err := c.repository.FindMember(id)
	if err != nil {
		return responses.ErrorResponse(ctx, 404, "member not found")
	}
	return responses.NewMemberResponse(ctx, model)
}

func (c *MemberController) UpdateMember(ctx *fiber.Ctx) error {
	/* === RUN VALIDATOR === */
	req := requests.UpdateMemberRequest{}
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
	birthDate, birthDateErr := utils.DateUtilYYYYMMDD(req.Birthdate)
	if birthDateErr != nil {
		return responses.ErrorValidationResponse(ctx, "invalid birthdate")
	}
	model := models.Member{
		ID:             parsedUUID,
		Phone:          req.Phone,
		Email:          req.Email,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		IdentityNumber: req.IdentityNumber,
		PlaceOfBirth:   req.PlaceOfBirth,
		Birthdate:      *birthDate,
		Gender:         req.Gender,
		Nationality:    req.Nationality,
		Address:        req.Address,
		PostalCode:     req.PostalCode,
		HomePhone:      req.HomePhone,
		OfficePhone:    req.OfficePhone,
		Education:      req.Education,
		Branches:       req.Branches,
	}
	err := c.repository.UpdateMember(&model)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return responses.ErrorValidationResponse(ctx, "member already exists")
		}
		return responses.ErrorBadRequest(ctx)
	}

	return responses.SuccessResponse(ctx, "member updated successfully")
}
