package routes

import (
	"golang-fiber-keycloak/handlers"
	"golang-fiber-keycloak/infrastructure/identity"
	registeruser "golang-fiber-keycloak/use_case/user/register-user"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome :)"))
	})

	group := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUserUseCase := registeruser.NewRegisterUserUseCase(identityManager)

	group.Post("/users", handlers.RegisterUserHandler(registerUserUseCase))
}
