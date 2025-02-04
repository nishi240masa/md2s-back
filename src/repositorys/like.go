package repositorys

import (
	"errors"
	"md2s/models"

	"gorm.io/gorm"
)


func CreateLike(newLike *models.Articlelike) error {
	
	// いいねを登録
	result := db.Create(newLike)

	// articlesテーブルのいいね数を更新
	db.Model(&models.Article{}).Where("id = ?", newLike.ArticleId).Update("like_count", gorm.Expr("like_count + ?", 1))
	
	if result.Error != nil {
		return result.Error
	}
	
	return nil
}

func DeleteLike(userId models.UUID, articleId int) error {

	// いいねを削除
	result := db.Where("user_id = ? AND article_id = ?", userId, articleId).Delete(&models.Articlelike{})

	// articlesテーブルのいいね数を更新
	db.Model(&models.Article{}).Where("id = ?", articleId).Update("like_count", gorm.Expr("like_count - ?", 1))
	
	if result.Error != nil {
		return result.Error
	}
	
	return nil

}

func GetLikes(userId models.UUID) ([]models.Article, error) {

	var articles []models.Article

	// いいねをした記事を取得
	result := db.Table("articles").Select("articles.*").Joins("JOIN articlelikes ON articles.id = articlelikes.article_id").Where("articlelikes.user_id = ?", userId).Find(&articles)

	if result.Error != nil {
		return nil, result.Error
	}

	
	return articles, nil

}

func GetLikesByArticleId(userId models.UUID, articleId int) (bool, error) {
	
	var like models.Articlelike
	result := db.Where("user_id = ? AND article_id = ?", userId, articleId).First(&like)
	
	// みつからなかった場合
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}else if result.Error != nil {
		return false, result.Error
	}
	
	return true, nil
}