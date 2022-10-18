package routes

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

// All defined routes to be expose for finance-service
func EndpointRoutes(app *fiber.App) {
	// all GET method for budgets endpoint
	app.Get("/health-check", controllers.HealthCheckHandler)

	app.Get("/budgets/:id", controllers.GetAllBudgetsController)
	app.Get("/budget/:budget_id", controllers.GetAllBudgetsController)

	// all POST method for budgets endpoint
	app.Post("/budgets", controllers.CreateBudgetHandler)
}
