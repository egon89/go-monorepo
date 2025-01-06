package routes

import (
	"golang-fiber-keycloak/handlers"
	"golang-fiber-keycloak/infrastructure/datastores"
	"golang-fiber-keycloak/infrastructure/identity"
	createproduct "golang-fiber-keycloak/use_case/product/create-product"
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

func InitProtectedRoutes(app *fiber.App) {
	group := app.Group("/api/v1")
	productDataStore := datastores.NewProductDataStore()
	createProductUseCase := createproduct.NewCreateProductUseCase(productDataStore)

	group.Post("/products", handlers.CreateProductHandler(createProductUseCase))
}
