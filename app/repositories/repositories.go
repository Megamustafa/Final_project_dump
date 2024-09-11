package repositories

import "aquaculture/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (models.Product, error)
	Create(productReq models.ProductRequest) (models.Product, error)
	Update(productReq models.ProductRequest, id string) (models.Product, error)
	Delete(id string) error
}

type AquacultureFarmsRepository interface {
	GetAll() ([]models.AquacultureFarms, error)
	GetByID(id string) (models.AquacultureFarms, error)
	Create(afReq models.AquacultureFarmsRequest) (models.AquacultureFarms, error)
	Update(afReq models.AquacultureFarmsRequest, id string) (models.AquacultureFarms, error)
	Delete(id string) error
}

type ProductTypeRepoistory interface {
	GetAll() ([]models.ProductType, error)
	GetByID(id string) (models.ProductType, error)
	Create(ptReq models.ProductTypeRequest) (models.ProductType, error)
	Update(ptReq models.ProductTypeRequest, id string) (models.ProductType, error)
	Delete(id string) error
}

type FarmTypeRepoistory interface {
	GetAll() ([]models.FarmType, error)
	GetByID(id string) (models.FarmType, error)
	Create(ftReq models.FarmTypeRequest) (models.FarmType, error)
	Update(ftReq models.FarmTypeRequest, id string) (models.FarmType, error)
	Delete(id string) error
}
