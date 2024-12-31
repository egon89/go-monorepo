package createproduct

import (
	"context"
	"golang-fiber-keycloak/domain/entities"
	"golang-fiber-keycloak/shared/utils"
	"golang-fiber-keycloak/use_case/product"
	"time"

	"github.com/google/uuid"
)

type createProductUseCase struct {
	dataStorage product.ProductDataStorage
}

func NewCreateProductUseCase(ds product.ProductDataStorage) *createProductUseCase {
	return &createProductUseCase{
		dataStorage: ds,
	}
}

func (uc *createProductUseCase) CreateProduct(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	var product = toProduct(input)

	err = uc.dataStorage.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		Product: product,
	}, nil
}

func toProduct(input CreateProductInput) *entities.Product {
	return &entities.Product{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      input.Name,
		Price:     input.Price,
	}
}
