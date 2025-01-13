package middlewares

import (
	"context"
	"golang-fiber-keycloak/infrastructure/identity"
	"golang-fiber-keycloak/shared/enums"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitFiberMiddlewares(
	app *fiber.App,
	initPublicRoutes func(app *fiber.App),
	initProtectedRoutes func(app *fiber.App),
) {
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		// get the request id that was added by requestid middleware
		requestId := c.Locals("requestid")

		// create a new context and add the requestId to it
		ctx := context.WithValue(context.Background(), enums.ContextKeyRequestId, requestId)
		c.SetUserContext(ctx)

		return c.Next()
	})

	initPublicRoutes(app)

	tokenRetrospector := identity.NewIdentityManager()
	app.Use(NewJwtMiddleware(tokenRetrospector))

	initProtectedRoutes(app)

	log.Println("fiber middlewares initialized")
}
