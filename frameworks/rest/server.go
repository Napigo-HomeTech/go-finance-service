package rest

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/middlewares"
	"github.com/Napigo/go-finance-service/frameworks/rest/routes"
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

var PORT = ":80"

func ListeningFunc() error {
	npglogger.Infof("Rest server listening on port %s", PORT)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorResponse,
	})

	CreateRestHook(app)
	middlewares.UserSessionMiddleware(app)

	routes.EndpointRoutes(app)

	app.Listen(PORT)
	app.Hooks().OnListen(ListeningFunc)
}
