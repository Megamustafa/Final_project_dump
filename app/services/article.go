package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
)

type ArticleService struct {
	repository repositories.ArticleRepository
}

func InitArticleService() ArticleService {
	return ArticleService{
		repository: &repositories.ArticleRepositoryImpl{},
	}
}

func (as *ArticleService) GetAll() ([]models.Article, error) {
	return as.repository.GetAll()
}

func (as *ArticleService) GetByID(id string) (models.Article, error) {
	return as.repository.GetByID(id)
}

func (as *ArticleService) Create(aReq models.ArticleRequest) (models.Article, error) {
	return as.repository.Create(aReq)
}

func (as *ArticleService) Update(aReq models.ArticleRequest, id string) (models.Article, error) {
	return as.repository.Update(aReq, id)
}

func (as *ArticleService) Delete(id string) error {
	return as.repository.Delete(id)
}
