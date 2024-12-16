package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/infra/database"
	"github.com/egon89/go-monorepo/user-product-api/internal/dto"
	"github.com/egon89/go-monorepo/user-product-api/internal/entity"
	"github.com/egon89/go-monorepo/user-product-api/pkg/helper"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder((r.Body)).Decode(&user)
	if err != nil {
		helper.ReturnErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		helper.ReturnErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		helper.ReturnErrorResponse(http.StatusInternalServerError, err, w)
		return
	}

	helper.CreatedResponse(u, w)
}
