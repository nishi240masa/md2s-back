package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)

func GetTags() ([]models.Tag, error) {
	return repositorys.GetTags()
}

func GetTag(id int) (*models.Tag, error) {
	return repositorys.GetTag(id)
}

func CreateTag(input []dto.CreateTagData)  ([]models.Tag, error) {
	var newTag models.Tag
	var newTags []models.Tag
	for _, tag := range input {

		newTag = models.Tag{
			Word: tag.Word,
		}

		err := repositorys.CreateTag(&newTag)
		if err != nil {
		return nil ,err
		}

		newTags = append(newTags, newTag)
	}

	return newTags, nil
}

func UpdateTag(id int, input dto.CreateTagData) error {
	
	tag, err := repositorys.GetTag(id)
	if err != nil {
		return err
	}

	tag.Word = input.Word

	return repositorys.UpdateTag(tag)
}

func DeleteTag(id int) error {
		return repositorys.DeleteTag(id)
}