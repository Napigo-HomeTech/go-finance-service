package routes

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(app *fiber.App) {
	// all GET method for budgets endpoint
	app.Get("/health-check", controllers.HealthCheckController)
}

// All defined routes to be expose for finance-service
func BudgetsRoutes(app *fiber.App) {
	app.Get("/budgets/:id", controllers.GetAllBudgetsController)

	// all POST method for budgets endpoint
	app.Post("/budgets", controllers.CreateBudgetController)
}
