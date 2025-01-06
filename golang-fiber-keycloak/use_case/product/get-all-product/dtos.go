package getallproduct

import "golang-fiber-keycloak/domain/entities"

type GetAllProductOutput struct {
	Products *[]entities.Product
}
