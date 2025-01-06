package handlers

import (
	"context"
	createproduct "golang-fiber-keycloak/use_case/product/create-product"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, input createproduct.CreateProductInput) (*createproduct.CreateProductOutput, error)
}

func CreateProductHandler(useCase CreateProductUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		var input = createproduct.CreateProductInput{}

		err := c.BodyParser(&input)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		output, err := useCase.CreateProduct(ctx, input)
		if err != nil {
			return err
		}

		log.Printf("product %s (%s) created", output.Product.Name, output.Product.Id)

		return c.Status(fiber.StatusCreated).JSON(output)
	}
}
