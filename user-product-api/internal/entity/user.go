package entity

import "github.com/egon89/go-monorepo/user-product-api/pkg/entity"

type User struct {
	Id       entity.Id `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}
