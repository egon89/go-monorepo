package getallproduct

import (
	"context"
	"golang-fiber-keycloak/use_case/product"
)

type getAllProductUseCase struct {
	dataStorage product.ProductDataStorage
}

func NewGetAllProductUseCase(ds product.ProductDataStorage) *getAllProductUseCase {
	return &getAllProductUseCase{
		dataStorage: ds,
	}
}

func (uc *getAllProductUseCase) GetAllProducts(ctx context.Context) *GetAllProductOutput {
	products := uc.dataStorage.GetAll()

	return &GetAllProductOutput{
		products: &products,
	}
}
