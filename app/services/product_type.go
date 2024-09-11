package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type ProductTypeService struct {
	repository repositories.ProductTypeRepoistory
}

func InitProductTypeService() ProductTypeService {
	return ProductTypeService{
		repository: &repositories.ProductTypeRepoistoryImpl{},
	}
}

func (pts *ProductTypeService) GetAll() ([]models.ProductType, error) {
	return pts.repository.GetAll()
}

func (pts *ProductTypeService) GetByID(id string) (models.ProductType, error) {
	return pts.repository.GetByID(id)
}

func (pts *ProductTypeService) Create(ptReq models.ProductTypeRequest) (models.ProductType, error) {
	return pts.repository.Create(ptReq)
}

func (pts *ProductTypeService) Update(ptReq models.ProductTypeRequest, id string) (models.ProductType, error) {
	return pts.repository.Update(ptReq, id)
}

func (pts *ProductTypeService) Delete(id string) error {
	return pts.repository.Delete(id)
}
