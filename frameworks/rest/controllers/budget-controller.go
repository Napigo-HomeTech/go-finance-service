package controllers

import (
	"github.com/Napigo/go-finance-service/internals/models"
	"github.com/Napigo/npgc"
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func GetAllBudgetsController(c *fiber.Ctx) error {
	sub, ok := c.UserContext().Value(npgc.UserSubKey).(string)

	npglogger.Infof("User Subject from Token : %s", sub)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "User id Not found")
	}

	dataList := []models.Budget{}

	for i := 0; i < 10; i++ {
		data := new(models.Budget)
		data.BudgetId = "1234ft"
		data.Revision = "JAN-2022"
		dataList = append(dataList, *data)
	}

	resp := npgc.RestResponse{Context: c, Payload: dataList, HTTPStatus: 200, Status: "success"}
	return resp.SendResponse()
}

func CreateBudgetController(c *fiber.Ctx) error {
	type ReqBodyCreateBudget struct {
		BudgetId string `json:"budget_id"`
		Name     string `json:"name"`
	}

	p := new(ReqBodyCreateBudget)

	if err := c.BodyParser(p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	resp := npgc.RestResponse{Context: c, Payload: "Created", HTTPStatus: 200, Status: "Success"}
	return resp.SendResponse()
}
