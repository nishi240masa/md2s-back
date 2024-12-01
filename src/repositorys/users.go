package repositorys

import (
	"errors"
	"fmt"
	"md2s/models"

	"gorm.io/gorm"
)

func GetUsers(sortOptions models.UserSortOptions) ([]models.User, error) {
	var users []models.User

	query := db.Order(fmt.Sprintf("%s %s", sortOptions.OrderBy, sortOptions.Order))
	result := query.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUserByGoogleID( googleID string) (*models.User, error) {
	var user models.User
	result := db.Where("google_id = ?", googleID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(newUser *models.User) error {
	result := db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}