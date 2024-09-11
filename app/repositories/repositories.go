package repositories

import "aquaculture/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Create(productReq models.ProductRequest) (models.Product, error)
	Update(productReq models.ProductRequest, id string) (models.Product, error)
	Delete(id string) error
}
