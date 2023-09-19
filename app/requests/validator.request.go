package requests

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidatorErrorData struct {
	FailedField string
	Tag         string
	Value       string
}

func phoneNumberOrEmpty(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	// Check if the field is empty
	if fieldValue == "" {
		return true
	}

	// Check if the field is numeric
	_, err := strconv.Atoi(fieldValue)
	if err != nil {
		return false
	}

	// Check if the field length is between 9 and 15 characters
	fieldLen := len(fieldValue)
	return fieldLen >= 9 && fieldLen <= 15
}

func emailOrEmpty(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	// Check if the field is empty
	if fieldValue == "" {
		return true
	}

	// Check if the field is a valid email address
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailPattern, fieldValue)

	// Check if the field length is between 5 and 50 characters
	fieldLen := len(fieldValue)
	return matched && fieldLen >= 5 && fieldLen <= 50
}

func NewValidatorRequest(ctx *fiber.Ctx, model interface{}) []*ValidatorErrorData {
	errors := []*ValidatorErrorData{}
	validate := validator.New()
	validate.RegisterValidation("phoneNumberOrEmpty", phoneNumberOrEmpty)
	validate.RegisterValidation("emailOrEmpty", emailOrEmpty)
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
