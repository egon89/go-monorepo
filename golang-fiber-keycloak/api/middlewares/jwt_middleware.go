package middlewares

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"golang-fiber-keycloak/shared/enums"

	"github.com/Nerzal/gocloak/v13"
	contribJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenRetrospector interface {
	RetrospectToken(context context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
}

func NewJwtMiddleware(tokenRetrospector TokenRetrospector) fiber.Handler {
	base64Str := viper.GetString("KEYCLOAK_REALM_PUBLIC_KEY")
	publicKey, err := parseKeycloakRSAPublicKey(base64Str)
	if err != nil {
		panic(err)
	}

	return contribJwt.New(contribJwt.Config{
		SigningKey: contribJwt.SigningKey{
			JWTAlg: contribJwt.RS256,
			Key:    publicKey,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			return processHandler(c, tokenRetrospector)
		},
	})
}

func parseKeycloakRSAPublicKey(base64Str string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString((base64Str))
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, nil
	}

	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}

func processHandler(c *fiber.Ctx, tokenRetrospector TokenRetrospector) error {
	jwtToken := c.Locals("user").(*golangJwt.Token)
	claims := jwtToken.Claims.(golangJwt.MapClaims)
	ctx := c.UserContext()

	contextWithClaims := context.WithValue(ctx, enums.ContextKeyClaims, claims)
	c.SetUserContext(contextWithClaims)

	rptResult, err := tokenRetrospector.RetrospectToken(ctx, jwtToken.Raw)
	if err != nil {
		panic(err)
	}

	if !*rptResult.Active {
		return c.Status(fiber.StatusUnauthorized).SendString("token is not active")
	}

	return c.Next()
}
