package rest

import "github.com/gofiber/fiber/v2"

type RestError struct {
	HTTPStatus int
	Message    string
}

func (re RestError) SendError(ctx *fiber.Ctx) {
	ctx.SendStatus(re.HTTPStatus)
	ctx.JSON(fiber.Map{
		"code":    re.HTTPStatus,
		"message": re.Message,
	})
}
