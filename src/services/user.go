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

func GetUserByJWT( jwtToken string) (*models.GetUser, error) {
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return nil, err
	}

	res, err :=	repositorys.GetUserByGoogleID( claims.Sub)
	if err != nil {
		return nil, err
	}

	var resUser models.GetUser
	if res.QiitaId != "" {
		resUser.Qiita_link = true
	}

	// userIdからlikesを取得
	likes, err := repositorys.GetLikes(res.ID)

	if err != nil {
		return nil, err
	}

	for _, like := range likes {
		resUser.Total_get_like_count += like.LikeCount
	}

	// userIdからarticlesを取得
	articles, err := repositorys.GetArticlesByUserId(res.ID)

	if err != nil {
		return nil, err
	}

	resUser.Total_posts_articles = len(articles)

	resUser.ID = res.ID
	resUser.Name = res.Name
	resUser.IconURL = res.IconURL

	return &resUser, nil
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
		QiitaId: "",
    }

    err = repositorys.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
