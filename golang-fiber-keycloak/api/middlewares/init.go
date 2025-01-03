package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiberMiddlewares(
	app *fiber.App,
	initPublicRoutes func(app *fiber.App),
) {
	app.Use(logger.New())

	initPublicRoutes(app)
}
