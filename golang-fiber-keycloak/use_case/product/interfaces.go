package product

import "golang-fiber-keycloak/domain/entities"

type ProductDataStorage interface {
	GetAll() []entities.Product
	Create(product *entities.Product) error
}
