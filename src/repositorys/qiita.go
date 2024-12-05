package repositorys

import (
	"md2s/models"
)


func AlignmentQiita(newData *models.User) error {

	result := db.Model(&models.User{}).Where("id = ?", newData.ID).Update("qiita_id", newData.QiitaId)

	if result.Error != nil {
		return result.Error
	}

	return nil

}