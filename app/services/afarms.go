package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type AquacultureFarmsService struct {
	repository repositories.AquacultureFarmsRepository
}

func InitAquacultureFarmsService() AquacultureFarmsService {
	return AquacultureFarmsService{
		repository: &repositories.AquacultureFarmsRepositoryImpl{},
	}
}

func (afs *AquacultureFarmsService) GetAll() ([]models.AquacultureFarms, error) {
	return afs.repository.GetAll()
}

func (afs *AquacultureFarmsService) GetByID(id string) (models.AquacultureFarms, error) {
	return afs.repository.GetByID(id)
}

func (afs *AquacultureFarmsService) Create(afReq models.AquacultureFarmsRequest) (models.AquacultureFarms, error) {
	return afs.repository.Create(afReq)
}

func (afs *AquacultureFarmsService) Update(afReq models.AquacultureFarmsRequest, id string) (models.AquacultureFarms, error) {
	return afs.repository.Update(afReq, id)
}

func (afs *AquacultureFarmsService) Delete(id string) error {
	return afs.repository.Delete(id)
}
