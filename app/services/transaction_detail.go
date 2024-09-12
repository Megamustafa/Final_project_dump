package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type TransactionDetailService struct {
	repository repositories.TransactionDetailRepository
}

func InitTransactionDetailService() TransactionDetailService {
	return TransactionDetailService{
		repository: &repositories.TransactionDetailRepositoryImpl{},
	}
}

func (tds *TransactionDetailService) GetAll() ([]models.TransactionDetail, error) {
	return tds.repository.GetAll()
}

func (tds *TransactionDetailService) GetByID(id string) (models.TransactionDetail, error) {
	return tds.repository.GetByID(id)
}

func (tds *TransactionDetailService) Create(tReq models.TransactionDetailRequest) (models.TransactionDetail, error) {
	return tds.repository.Create(tReq)
}

func (tds *TransactionDetailService) Update(tReq models.TransactionDetailRequest, id string) (models.TransactionDetail, error) {
	return tds.repository.Update(tReq, id)
}
