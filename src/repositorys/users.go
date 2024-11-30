package repositorys

import (
	"md2s/models"
)

func GetUserById(id uint) (*models.User, error) {

	var  user models.User

	result := db.Where("id= ? " , id).First(&user)

	if result.Error != nil {
		return  nil,result.Error
	}
	return &user, nil

}

func CreateUser(user models.UserCreate) (*models.User, error) {
	
	var newUser models.User

	newUser.Name = user.Name
	newUser.IconURL = user.IconURL
	newUser.GoogleId = user.GoogleId

	result := db.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func DeleteUser(id models.UUID) error {
	result := db.Where("id= ? ", id).Delete(&models.User{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

