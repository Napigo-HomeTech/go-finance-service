package rest

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/routes"
	"github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

var PORT = ":80"

func ListeningFunc() error {
	logger.Infof("Rest server listening on port %s", PORT)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorResponse,
	})

	CreateRestHook(app)
	routes.EndpointRoutes(app)

	app.Listen(PORT)
	app.Hooks().OnListen(ListeningFunc)
}
