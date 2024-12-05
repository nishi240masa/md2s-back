package services

import (
	"errors"
	"fmt"
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)

func GetUsers( sortOptions models.UserSortOptions) ([]models.User, error) {
	return repositorys.GetUsers( sortOptions)
}

func GetUserByJWT( jwtToken string) (*models.User, error) {
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return nil, err
	}

	return repositorys.GetUserByGoogleID( claims.Email) // Google IDとしてEmailを利用
}

func CreateUser(jwtToken string, input dto.CreateUserData) (*models.User, error) {
    claims, err := VerifyGoogleToken(jwtToken)
    if err != nil {
        return nil, err
    }

    // inputデータとclaimsが一致するか確認
	if claims.Sub != input.GoogleId {

		fmt.Println("sub",claims.Sub)
		fmt.Println("inputID",input.GoogleId)


		return nil, errors.New("invalid Google ID")
	}

	// 既に登録されているか確認
	_, err = repositorys.GetUserByGoogleID(input.GoogleId)
	if err == nil {
		return nil, errors.New("user already exists")
	} 

    newUser := &models.User{
        Name:     input.Name,
        IconURL:  input.IconURL,
        GoogleId: input.GoogleId,
    }

    err = repositorys.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
