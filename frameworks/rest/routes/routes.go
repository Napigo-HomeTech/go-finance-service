package routes

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoites(app *fiber.App) {
	// all GET method for budgets endpoint
	app.Get("/health-check", controllers.HealthCheckHandler)
}

// All defined routes to be expose for finance-service
func BudgetsRoutes(app *fiber.App) {
	app.Get("/budgets/:id", controllers.GetAllBudgetsController)
	app.Get("/budget/:budget_id", controllers.GetAllBudgetsController)

	// all POST method for budgets endpoint
	app.Post("/budgets", controllers.CreateBudgetHandler)
}
