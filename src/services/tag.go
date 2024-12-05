package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)

func GetTags() ([]string, error) {
	return repositorys.GetTags()
}

func CreateTag(input dto.CreateTagData) error {

	newTag := models.Tag{
		Word: input.Word,
	}

	return repositorys.CreateTag(&newTag)
}