package rest

import (
	"os"

	"github.com/Napigo/go-finance-service/frameworks/rest/routes"
	"github.com/Napigo/npgcommon/rest"
	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func ListeningFunc() error {
	service_name := os.Getenv("SERVICE_NAME")
	port := os.Getenv("SERVICE_PORT")
	npglogger.Infof("Server %s listening on port %s", service_name, port)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	port := os.Getenv("SERVICE_PORT")
	app := fiber.New(fiber.Config{
		ErrorHandler: rest.DefaultErrorResponse,
	})

	rest.CreateRestHook(app)
	rest.UserSubMiddleware(app)

	routes.BudgetsRoutes(app)
	routes.BudgetsRoutes(app)

	app.Hooks().OnListen(ListeningFunc)
	app.Listen(port)
}
