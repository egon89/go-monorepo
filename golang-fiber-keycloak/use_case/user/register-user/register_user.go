package registeruser

import (
	"context"
	"golang-fiber-keycloak/shared/utils"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

type registerUserUseCase struct {
	identityManager identityManager
}

func NewRegisterUserUseCase(im identityManager) *registerUserUseCase {
	return &registerUserUseCase{
		identityManager: im,
	}
}

func (uc *registerUserUseCase) Register(ctx context.Context, input RegisterUserInput) (*RegisterUserOutput, error) {
	err := utils.ValidateStruct(input)
	if err != nil {
		return nil, err
	}

	var user = toUserKeycloak(input)

	userCreated, err := uc.identityManager.CreateUser(ctx, user, input.Password, "viewer")
	if err != nil {
		return nil, err
	}

	return &RegisterUserOutput{
		User: userCreated,
	}, nil

}

func toUserKeycloak(input RegisterUserInput) gocloak.User {
	user := gocloak.User{
		Username:      gocloak.StringP(input.Username),
		FirstName:     gocloak.StringP(input.FirstName),
		LastName:      gocloak.StringP(input.LastName),
		Email:         gocloak.StringP(input.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
	}

	if strings.TrimSpace(input.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{input.MobileNumber}
	}

	return user
}
