package models

import "github.com/go-playground/validator/v10"

type TransactionRequest struct {
	UserID        uint   `json:"user_id"`
	TotalAmount   uint   `json:"total_amount"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
}

type TransactionDetailRequest struct {
	TransactionID uint `json:"transaction_id"`
	ProductID     uint `json:"product_id"`
	Quantity      uint `json:"quantity"`
	Amount        uint `json:"amount"`
}

func (t *TransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	return err
}

func (td *TransactionDetailRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(td)

	return err
}
