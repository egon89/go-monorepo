package entity

import (
	"errors"
	"time"

	"github.com/egon89/go-monorepo/user-product-api/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	Id        entity.Id `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		Id:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Id.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseId(p.Id.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
