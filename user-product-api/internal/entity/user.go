package entity

import (
	"github.com/egon89/go-monorepo/user-product-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       entity.Id `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err

	}

	return &User{
		Id:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}
