package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductRequest struct {
	ID          uint           `json:"id" validate:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	TypeID      uint           `json:"type_ID" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Price       int            `json:"price" validate:"required"`
	ProductType ProductType    `json:"product_type"`
}

func (p *ProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(p)

	return err
}
