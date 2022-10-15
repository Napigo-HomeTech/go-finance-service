package rest

import "github.com/gofiber/fiber/v2"

func DefaultErrorResponse(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return ctx.JSON(fiber.Map{
		"code":    code,
		"status":  "error",
		"message": message,
	})
}
