package rest

import (
	"github.com/Napigo/go-finance-service/frameworks/rest/routes"
	"github.com/Napigo/npgc"

	"github.com/Napigo/npglogger"
	"github.com/gofiber/fiber/v2"
)

func ListeningFunc() error {
	config := npgc.Config
	npglogger.Infof("Server %s listening on port %s", config.ServiceName, config.Port)
	return nil
}

type RestServer struct{}

func (rs RestServer) Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: npgc.DefaultErrorResponse,
	})

	npgc.CreateRestHook(app)
	npgc.AuthVerify(app)

	routes.HealthCheckRoutes(app)
	routes.BudgetsRoutes(app)
	routes.BudgetsRoutes(app)

	app.Hooks().OnListen(ListeningFunc)

	port := npgc.Config.Port
	app.Listen(port)
}
