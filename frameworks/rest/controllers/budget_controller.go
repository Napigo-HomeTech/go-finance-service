package controllers

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/restutils"
	"github.com/Napigo/go-finance-service/internals/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllBudgetsController(c *fiber.Ctx) error {
	userId := c.Query("user_id")

	if len(userId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "user_id is not undefined")
	}

	dataList := []models.Budget{}

	for i := 0; i < 10; i++ {
		data := new(models.Budget)
		data.BudgetId = "1234ft"
		data.Revision = "JAN-2022"
		dataList = append(dataList, *data)
	}

	resp := restutils.RestResponse{Context: c, Payload: dataList}
	return resp.SendResponse()
}

type ReqBodyCreateBudget struct {
	BudgetId string `json:"budget_id"`
	Name     string `json:"name"`
}

func CreateBudgetHandler(c *fiber.Ctx) error {
	p := new(ReqBodyCreateBudget)

	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	resp := restutils.RestResponse{Context: c, Payload: "Created", HTTPStatus: 200, Status: "Success"}
	return resp.SendResponse()
}
