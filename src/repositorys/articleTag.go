package repositorys

import "md2s/models"

// タグの登録
func CreateArticleTag(newArticleTag *models.Articletagrelations) error {


	result := db.Create(newArticleTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 記事IDからタグを取得
func GetArticleTagByArticleID(articleID int) ([]models.Articletagrelations, error) {
	var articleTag []models.Articletagrelations
	result := db.Where("article_id = ?", articleID).Find(&articleTag)
	if result.Error != nil {
		return nil, result.Error
	}
	return articleTag, nil
}

// タグIDから記事を取得
func GetArticleTagByTagID(tagID int) ([]models.Articletagrelations, error) {
	var articleTag []models.Articletagrelations
	result := db.Where("tag_id = ?", tagID).Find(&articleTag)
	if result.Error != nil {
		return nil, result.Error
	}
	return articleTag, nil
}

// 記事IDからタグを削除
func DeleteArticleTagByArticleID(articleID int) error {
	result := db.Where("article_id = ?", articleID).Delete(&models.Articletagrelations{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// タグIDから記事を削除
func DeleteArticleTagByTagID(tagID int) error {
	result := db.Where("tag_id = ?", tagID).Delete(&models.Articletagrelations{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}