package routes

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

// All defined routes to be expose for finance-service
func EndpointRoutes(app *fiber.App) {
	route := app.Group("api/v1")

	// all GET method for budgets endpoint
	route.Get("/health-check", controllers.HealthCheckHandler)
	route.Get("/budgets", controllers.GetAllBudgetsController)
	route.Get("/budgets/:id", controllers.GetAllBudgetsController)

	// all POST method for budgets endpoint
	route.Post("/budgets", controllers.CreateBudgetHandler)
}
