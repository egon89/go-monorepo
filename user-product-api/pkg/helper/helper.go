package helper

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/pkg/dto"
)

func ReturnErrorResponse(httpStatus int, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	error := dto.ErrorResponse{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(error)
}

func CreatedResponse(output interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
