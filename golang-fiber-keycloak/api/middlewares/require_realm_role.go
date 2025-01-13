package middlewares

import (
	"golang-fiber-keycloak/shared/enums"
	"golang-fiber-keycloak/shared/jwt"

	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
)

func NewRequireRealmRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		claims := ctx.Value(enums.ContextKeyClaims).(golangJwt.MapClaims)
		jwtHelper := jwt.NewJwtHelper(claims)
		if !jwtHelper.HasRole(role) {
			return c.Status(fiber.StatusUnauthorized).SendString("role authorization failed")
		}

		return c.Next()
	}
}
