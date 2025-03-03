package repositorys

import (
	"md2s/dto"
	"md2s/models"
)


func GetArticles(query dto.GetArticlesData) ([]models.Article, error) {
	// 記事情報を取得
	var articles []models.Article


	// 記事のuser_idからユーザーiconとnameも取得
	// publicの記事のみ取得
	// offsetが0の場合はランダムに取得

	if query.Offset == 0 {
		
		result := db.
		Joins("JOIN users ON users.id = articles.user_id").
		Where("articles.public = ?", true).
		Select("articles.*, users.icon_url, users.name").
		Order("RANDOM()").
		Limit(query.Limit).
		Find(&articles)

		// ０件の場合
		if result.Error  == nil && len(articles) == 0 {
			return articles, nil
		}

		if result.Error != nil {
			return nil, result.Error
		}
		return articles, nil
	}
	result := db.
    Joins("JOIN users ON users.id = articles.user_id").
	Where("articles.public = ?", true).
	Select("articles.*, users.icon_url, users.name").
    Limit(query.Limit).
    Offset(query.Offset).
    Find(&articles)


	// ０件の場合
	if result.Error  == nil && len(articles) == 0 {
		return articles, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}
	
	return articles, nil

}

func GetArticle(id int) (*models.Article, error) {
	var article models.Article
	result := db.
	Joins("JOIN users ON users.id = articles.user_id").
	Select("articles.*, users.icon_url, users.name").
	Where("articles.id = ?", id).First(&article)


	if result.Error != nil {
		return nil, result.Error
	}
	return &article, nil
}
func GetArticlesByUserId(user_id models.UUID) ([]models.Article, error) {
	var articles []models.Article
	result := db.
	Joins("JOIN users ON users.id = articles.user_id").
	Select("articles.*, users.icon_url, users.name").
	Where("user_id = ?", user_id).
	Find(&articles)


		// ０件の場合
		if result.Error  == nil && len(articles) == 0 {
			return articles, nil
		}

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func CreateArticle(newArticle *models.Articles) (int, error) {


	
	result := db.Create(newArticle)
	if result.Error != nil {
		return 0, result.Error
	}

	return newArticle.ID, nil
}

func SearchArticles(input dto.SearchArticlesData) ([]models.Article, error) {
	var articles []models.Article
	result := db.
    Table("articles").
    Joins("JOIN articleTagRelations ON articleTagRelations.article_id = articles.id").
    Joins("JOIN tags ON tags.id = articleTagRelations.tag_id").
    Joins("JOIN users ON users.id = articles.user_id").
    Select("DISTINCT articles.*, users.icon_url AS icon_url, users.name AS user_name").
	Where("articles.public = ?", true).
    Where("articles.title LIKE ? OR tags.word LIKE ?", "%"+input.Keyword+"%", "%"+input.Keyword+"%").
    Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func UpdateArticle(article *models.Articles) error {
	result := db.Save(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteArticle(id int) error {
	result := db.Where("id = ?", id).Delete(&models.Article{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}