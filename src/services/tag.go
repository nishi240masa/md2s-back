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

func CreateTag(input dto.CreateTagData) error {

	newTag := models.Tag{
		Word: input.Word,
	}

	return repositorys.CreateTag(&newTag)
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