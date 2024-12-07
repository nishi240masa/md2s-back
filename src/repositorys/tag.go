package repositorys

import "md2s/models"


func GetTags() ([]models.Tag, error) {

	// 全部のデータを取得
	var tags []models.Tag
	result := db.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}

// GetTag は指定したIDのタグを取得する
func GetTag(id int) (*models.Tag, error) {
	var tag models.Tag
	result := db.Where("id = ?", id).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func SearchTags (keyword string) ([]models.Tag, error) {
	var tags []models.Tag
	result := db.Where("word LIKE ?", "%" + keyword + "%").Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}


func CreateTag(newTag *models.Tag) error {

	result := db.Create(newTag)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func UpdateTag(newTag *models.Tag) error {
	
	result := db.Save(newTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTag(id int) error {
	result := db.Where("id = ?", id).Delete(&models.Tag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

