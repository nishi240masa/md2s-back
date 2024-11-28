package repositorys

import "md2s/models"

func GetMemberById(id uint) (*models.Member, error) {

	var  member models.Member

	result := db.Where("id= ? " , id).First(&member)

	if result.Error != nil {
		return  nil,result.Error
	}
	return &member, nil

}