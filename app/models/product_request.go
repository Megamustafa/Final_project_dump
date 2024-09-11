package models

import (
	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	ProductTypeID uint   `json:"type_ID" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Price         int    `json:"price" validate:"required"`
}

func (p *ProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	return err
}
