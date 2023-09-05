package responses

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(c *fiber.Ctx, data *Response) error {
	return c.Status(data.Status).JSON(data)
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	res := Response{
		Status:  fiber.StatusOK,
		Message: http.StatusText(fiber.StatusOK),
		Data:    data,
	}
	return NewResponse(c, &res)
}

func Resource(ctx *fiber.Ctx, model interface{}, res interface{}) error {
	data := &res
	jsonMarshal, err := json.Marshal(model)
	if err != nil {
		return err
	}
	errUnmarshal := json.Unmarshal(jsonMarshal, data)
	if errUnmarshal != nil {
		return errUnmarshal
	}
	return SuccessResponse(ctx, res)
}

func Collections(ctx *fiber.Ctx, model []interface{}, res *[]interface{}) error {
	jsonMarshal, err := json.Marshal(&model)
	if err != nil {
		return err
	}
	errUnmarshal := json.Unmarshal(jsonMarshal, res)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return SuccessResponse(ctx, res)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	res := Response{
		Status:  statusCode,
		Message: message,
		Data:    nil,
	}
	return NewResponse(c, &res)
}

func ErrorInternal(c *fiber.Ctx) error {
	res := Response{
		Status:  fiber.StatusInternalServerError,
		Message: http.StatusText(fiber.StatusInternalServerError),
		Data:    nil,
	}
	return NewResponse(c, &res)
}

func ErrorBadRequest(c *fiber.Ctx) error {
	res := Response{
		Status:  fiber.StatusBadRequest,
		Message: http.StatusText(fiber.StatusBadRequest),
		Data:    nil,
	}
	return NewResponse(c, &res)
}

func ErrorValidationResponse(c *fiber.Ctx, data interface{}) error {
	res := Response{
		Status:  fiber.StatusUnprocessableEntity,
		Message: http.StatusText(fiber.StatusUnprocessableEntity),
		Data:    data,
	}
	return NewResponse(c, &res)
}

func ErrorUnauthorized(c *fiber.Ctx) error {
	return ErrorResponse(
		c,
		fiber.StatusUnauthorized,
		http.StatusText(fiber.StatusUnauthorized),
	)
}

func ErrorForbidden(c *fiber.Ctx) error {
	return ErrorResponse(
		c,
		fiber.StatusForbidden,
		http.StatusText(fiber.StatusForbidden),
	)
}
