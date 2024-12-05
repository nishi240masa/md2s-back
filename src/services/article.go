package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)


func GetArticles( quary dto.GetArticlesData) ([]models.Article, error) {

	return  repositorys.GetArticles(quary)
}

func GetArticle(id int) (*models.Article, error) {
	return repositorys.GetArticle(id)
}