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

func (tds *TransactionDetailService) Create(tdReq models.TransactionDetailRequest) (models.TransactionDetail, error) {
	return tds.repository.Create(tdReq)
}

func (tds *TransactionDetailService) Update(tdReq models.TransactionDetailRequest, id string) (models.TransactionDetail, error) {
	return tds.repository.Update(tdReq, id)
}
