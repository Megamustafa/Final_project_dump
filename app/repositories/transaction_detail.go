package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type TransactionDetailRepositoryImpl struct{}

func InitTransactionDetailRepository() TransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{}
}

func (tdr *TransactionDetailRepositoryImpl) GetAll() ([]models.TransactionDetail, error) {
	var transactionDetails []models.TransactionDetail

	if err := database.DB.Find(&transactionDetails).Error; err != nil {
		return []models.TransactionDetail{}, err
	}

	return transactionDetails, nil
}

func (tdr *TransactionDetailRepositoryImpl) GetByID(id string) (models.TransactionDetail, error) {
	var transactionDetail models.TransactionDetail

	if err := database.DB.First(&transactionDetail, "id = ?", id).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}

func (tdr *TransactionDetailRepositoryImpl) Create(tdReq models.TransactionDetailRequest) (models.TransactionDetail, error) {
	var transactionDetail models.TransactionDetail = models.TransactionDetail{
		TransactionID: tdReq.TransactionID,
		ProductID:     tdReq.ProductID,
		Quantity:      tdReq.Quantity,
		Amount:        tdReq.Amount,
	}

	result := database.DB.Create(&transactionDetail)

	if err := result.Error; err != nil {
		return models.TransactionDetail{}, err
	}

	if err := result.Last(&transactionDetail).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}

func (tdr *TransactionDetailRepositoryImpl) Update(tdReq models.TransactionDetailRequest, id string) (models.TransactionDetail, error) {
	transactionDetail, err := tdr.GetByID(id)

	if err != nil {
		return models.TransactionDetail{}, err
	}

	transactionDetail.TransactionID = tdReq.TransactionID
	transactionDetail.ProductID = tdReq.ProductID
	transactionDetail.Quantity = tdReq.Quantity
	transactionDetail.Amount = tdReq.Amount

	if err := database.DB.Save(&transactionDetail).Error; err != nil {
		return models.TransactionDetail{}, err
	}

	return transactionDetail, nil
}
