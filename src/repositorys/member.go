package repositorys

import "md2s/models"

func GetMemberById(id uint) (*models.User, error) {

	var  member models.User

	result := db.Where("id= ? " , id).First(&member)

	if result.Error != nil {
		return  nil,result.Error
	}
	return &member, nil

}