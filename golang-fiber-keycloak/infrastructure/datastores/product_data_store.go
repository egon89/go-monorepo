package datastores

import (
	"golang-fiber-keycloak/domain/entities"
	"slices"
	"sync"

	"github.com/google/uuid"
)

type productDataStore struct {
	products map[uuid.UUID]entities.Product
	sync.Mutex
}

func NewProductDataStore() *productDataStore {
	return &productDataStore{
		products: make(map[uuid.UUID]entities.Product),
	}
}

func (ds *productDataStore) Create(product *entities.Product) error {
	ds.Lock()
	ds.products[product.Id] = *product
	ds.Unlock()

	return nil
}

func (ds *productDataStore) GetAll() []entities.Product {
	all := make([]entities.Product, 0, len(ds.products))
	for _, value := range ds.products {
		all = append(all, value)
	}

	slices.SortFunc(all, func(i, j entities.Product) int {
		return i.CreatedAt.UTC().Compare(j.CreatedAt.UTC())
	})

	return all
}
