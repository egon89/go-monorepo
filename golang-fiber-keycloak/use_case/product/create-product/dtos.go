package createproduct

import "golang-fiber-keycloak/domain/entities"

type CreateProductInput struct {
	Name  string  `validate:"required,min=3,max=15"`
	Price float32 `validate:"required"`
}

type CreateProductOutput struct {
	Product *entities.Product
}
