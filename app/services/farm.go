package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type FarmService struct {
	repository repositories.FarmRepository
}

func InitFarmService() FarmService {
	return FarmService{
		repository: &repositories.FarmRepositoryImpl{},
	}
}

func (fs *FarmService) GetAll() ([]models.Farm, error) {
	return fs.repository.GetAll()
}

func (fs *FarmService) GetByID(id string) (models.Farm, error) {
	return fs.repository.GetByID(id)
}

func (fs *FarmService) Create(farmReq models.FarmRequest) (models.Farm, error) {
	return fs.repository.Create(farmReq)
}

func (fs *FarmService) Update(farmReq models.FarmRequest, id string) (models.Farm, error) {
	return fs.repository.Update(farmReq, id)
}

func (fs *FarmService) Delete(id string) error {
	return fs.repository.Delete(id)
}
