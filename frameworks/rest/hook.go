package rest

import (
	"github.com/Napigo/go-finance-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func CreateRestHook(app *fiber.App) {
	app.Hooks().OnGroup(func(g fiber.Group) error {
		logger.Info("Group route added, path :" + g.Prefix)
		return nil
	})

	app.Hooks().OnRoute(func(r fiber.Route) error {
		logger.Info("Route added, path : " + r.Path + " method: " + r.Method)
		return nil
	})
}
