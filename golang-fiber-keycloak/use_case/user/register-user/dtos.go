package registeruser

import "github.com/Nerzal/gocloak/v13"

type RegisterUserInput struct {
	Username     string `validate:"required,min=3,max=15"`
	Password     string `validate:"required"`
	FirstName    string `validate:"min=1,max=30"`
	LastName     string `validate:"min=1,max=30"`
	Email        string `validate:"required,email"`
	MobileNumber string
}

type RegisterUserOutput struct {
	User *gocloak.User
}
