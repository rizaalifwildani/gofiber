package requests

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidatorErrorData struct {
	FailedField string
	Tag         string
	Value       string
}

func NewValidatorRequest(ctx *fiber.Ctx, model interface{}) []*ValidatorErrorData {
	errors := []*ValidatorErrorData{}
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidatorErrorData
			element.FailedField = strings.ToLower(err.Field())
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
