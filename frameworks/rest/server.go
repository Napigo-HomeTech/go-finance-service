package rest

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/routes"
	"github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type RestServer struct{}

func (rs RestServer) Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorResponse,
	})

	routes.EndpointRoutes(app)

	app.Listen(":4000")
	logger.Info("Rest server listening on port :4000")
}
