package getallproduct

import "golang-fiber-keycloak/domain/entities"

type GetAllProductOutput struct {
	products *[]entities.Product
}