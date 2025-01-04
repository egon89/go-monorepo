package handlers

import (
	"context"
	registeruser "golang-fiber-keycloak/use_case/user/register-user"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type RegisterUserUseCase interface {
	// this should follow the use case method signature
	Register(ctx context.Context, request registeruser.RegisterUserInput) (*registeruser.RegisterUserOutput, error)
}

func RegisterUserHandler(useCase RegisterUserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		var input = registeruser.RegisterUserInput{}

		err := c.BodyParser(&input)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		output, err := useCase.Register(ctx, input)
		if err != nil {
			return errors.Wrapf(err, "error creating user %v", input.Username)
		}

		return c.Status(fiber.StatusCreated).JSON(output)
	}
}
