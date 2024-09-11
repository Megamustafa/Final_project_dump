package repositories

import (
	"aquaculture/database"
	"aquaculture/models"
)

type ProductTypeRepoistoryImpl struct{}

func InitProductTypeRepoistory() ProductTypeRepoistory {
	return &ProductTypeRepoistoryImpl{}
}

func (ptr *ProductTypeRepoistoryImpl) GetAll() ([]models.ProductType, error) {
	var productTypes []models.ProductType

	if err := database.DB.Find(&productTypes).Error; err != nil {
		return []models.ProductType{}, err
	}

	return productTypes, nil
}

func (ptr *ProductTypeRepoistoryImpl) GetByID(id string) (models.ProductType, error) {
	var productType models.ProductType

	if err := database.DB.First(&productType, "id = ?", id).Error; err != nil {
		return models.ProductType{}, err
	}

	return productType, nil
}

func (ptr *ProductTypeRepoistoryImpl) Create(ptReq models.ProductTypeRequest) (models.ProductType, error) {
	var productType models.ProductType = models.ProductType{
		Name: ptReq.Name,
	}

	result := database.DB.Create(&productType)

	if err := result.Error; err != nil {
		return models.ProductType{}, err
	}

	if err := result.Last(&productType).Error; err != nil {
		return models.ProductType{}, err
	}

	return productType, nil
}

func (ptr *ProductTypeRepoistoryImpl) Update(ptReq models.ProductTypeRequest, id string) (models.ProductType, error) {
	productType, err := ptr.GetByID(id)

	if err != nil {
		return models.ProductType{}, err
	}

	productType.Name = ptReq.Name

	if err := database.DB.Save(&productType).Error; err != nil {
		return models.ProductType{}, err
	}

	return productType, nil
}

func (ptr *ProductTypeRepoistoryImpl) Delete(id string) error {
	productType, err := ptr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&productType).Error; err != nil {
		return err
	}

	return nil
}
