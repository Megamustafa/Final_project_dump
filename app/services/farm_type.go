package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type FarmTypeService struct {
	repository repositories.FarmTypeRepository
}

func InitFarmTypeService() FarmTypeService {
	return FarmTypeService{
		repository: &repositories.FarmTypeRepositoryImpl{},
	}
}

func (fts *FarmTypeService) GetAll() ([]models.FarmType, error) {
	return fts.repository.GetAll()
}

func (fts *FarmTypeService) GetByID(id string) (models.FarmType, error) {
	return fts.repository.GetByID(id)
}

func (fts *FarmTypeService) Create(ftReq models.FarmTypeRequest) (models.FarmType, error) {
	return fts.repository.Create(ftReq)
}

func (fts *FarmTypeService) Update(ftReq models.FarmTypeRequest, id string) (models.FarmType, error) {
	return fts.repository.Update(ftReq, id)
}

func (fts *FarmTypeService) Delete(id string) error {
	return fts.repository.Delete(id)
}
