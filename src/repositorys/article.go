package repositorys

import (
	"md2s/dto"
	"md2s/models"
)


func GetArticles(query dto.GetArticlesData) ([]models.Article, error) {
	// 記事情報を取得
	var articles []models.Article


	// 記事のuser_idからユーザーiconとnameも取得
	result := db.
    Joins("LEFT JOIN articleTagRelations ON articleTagRelations.article_id = articles.id").
    Joins("LEFT JOIN tags ON tags.id = articleTagRelations.tag_id").
    Joins("JOIN users ON users.id = articles.user_id").
    Select("articles.*, COALESCE(tags.word, '') AS tag, users.name AS user_name, users.icon_url AS user_icon").
    Limit(query.Limit).
    Offset(query.Offset).
    Find(&articles)
	


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

func CreateArticle(newArticle *models.CreateArticle) (int, error) {
	
	result := db.Create(newArticle)
	if result.Error != nil {
		return 0, result.Error
	}

	return newArticle.ID, nil
}

func SearchArticles(input dto.SearchArticlesData) ([]models.Article, error) {
	var articles []models.Article
	result := db.
    Joins("JOIN articleTagRelations ON articleTagRelations.article_id = articles.id").
    Joins("JOIN tags ON tags.id = articleTagRelations.tag_id").
    Where("articles.title LIKE ? OR tags.word LIKE ?", "%"+input.Keyword+"%", "%"+input.Keyword+"%").
    Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func UpdateArticle(article *models.CreateArticle) error {
	result := db.Save(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}