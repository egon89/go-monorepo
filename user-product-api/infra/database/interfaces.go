package database

import "github.com/egon89/go-monorepo/user-product-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
}

type ProductInterface interface {
	Create(product *entity.Product) error
}
