package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type TransactionService struct {
	repository repositories.TransactionRepository
}

func InitTransactionService() TransactionService {
	return TransactionService{
		repository: &repositories.TransactionRepositoryImpl{},
	}
}

func (ts *TransactionService) GetAll() ([]models.Transaction, error) {
	return ts.repository.GetAll()
}

func (ts *TransactionService) GetByID(id string) (models.Transaction, error) {
	return ts.repository.GetByID(id)
}

func (ts *TransactionService) Create(tReq models.TransactionRequest) (models.Transaction, error) {
	return ts.repository.Create(tReq)
}

func (ts *TransactionService) Update(tReq models.TransactionRequest, id string) (models.Transaction, error) {
	return ts.repository.Update(tReq, id)
}
