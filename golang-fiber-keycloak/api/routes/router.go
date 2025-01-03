package routes

import "github.com/gofiber/fiber/v2"

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome :)"))
	})
}
