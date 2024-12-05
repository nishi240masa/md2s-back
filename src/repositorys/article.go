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

func CreateArticle(newArticle *models.Article) (int, error) {
	
	result := db.Create(newArticle)
	if result.Error != nil {
		return 0, result.Error
	}

	return newArticle.ID, nil
}

func UpdateArticle(article *models.Article) error {
	result := db.Save(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}