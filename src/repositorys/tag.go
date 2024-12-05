package repositorys

import "md2s/models"


func GetTags() ([]string, error) {
	var tags []string
	result := db.Table("tags").Pluck("name", &tags)
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