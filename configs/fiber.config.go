package configs

import (
	"errors"
	"net/http"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"github.com/gofiber/fiber/v2"
)

func InitFiberConfig() fiber.Config {
	return fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = responses.ErrorResponse(ctx, code, http.StatusText(code))
			if err != nil {
				// In case the SendFile fails
				return responses.ErrorInternal(ctx)
			}

			// Return from handler
			return nil
		},
	}
}
