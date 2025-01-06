package handlers

import (
	"context"
	getallproduct "golang-fiber-keycloak/use_case/product/get-all-product"

	"github.com/gofiber/fiber/v2"
)

type GetAllProductUseCase interface {
	GetAllProducts(ctx context.Context) *getallproduct.GetAllProductOutput
}

func GetAllProductHandler(useCase GetAllProductUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		products := useCase.GetAllProducts(ctx)

		return c.Status(fiber.StatusOK).JSON(products)
	}
}
