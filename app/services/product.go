package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type ProductService struct {
	repository repositories.ProductRepository
}

func InitProductService() ProductService {
	return ProductService{
		repository: &repositories.ProductRepositoryImpl{},
	}
}

func (ps *ProductService) GetAll() ([]models.Product, error) {
	return ps.repository.GetAll()
}

func (ps *ProductService) GetByID(id string) (models.Product, error) {
	return ps.repository.GetByID(id)
}

func (ps *ProductService) Create(productReq models.ProductRequest) (models.Product, error) {
	return ps.repository.Create(productReq)
}

func (ps *ProductService) Update(productReq models.ProductRequest, id string) (models.Product, error) {
	return ps.repository.Update(productReq, id)
}

func (ps *ProductService) Delete(id string) error {
	return ps.repository.Delete(id)
}
