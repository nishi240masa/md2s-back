package repositorys

import (
	"md2s/dto"
	"md2s/models"
)


func GetArticles(quary dto.GetArticlesData) ([]models.Article, error) {
	var articles []models.Article

	result := db.Limit(quary.Limit).Offset(quary.Offset).Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil

}

func GetArticle(id int) (*models.Article, error) {
	var article models.Article
	result := db.Where("id = ?", id).First(&article)
	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}