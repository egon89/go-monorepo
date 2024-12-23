package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-monorepo/user-product-api/infra/database"
	"github.com/egon89/go-monorepo/user-product-api/internal/dto"
	"github.com/egon89/go-monorepo/user-product-api/internal/entity"
	"github.com/egon89/go-monorepo/user-product-api/pkg/helper"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		helper.ReturnErrorResponse(400, err, w)
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		helper.ReturnErrorResponse(400, err, w)
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		helper.ReturnErrorResponse(500, err, w)
	}

	helper.CreatedResponse(p, w)
}
