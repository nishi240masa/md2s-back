package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)

func CreateLike(jwtToken string, input dto.Like) (error) {

	// 本人確認
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return err
	}

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(claims.Sub)
	if err != nil {
		return err
	}

	newLike := models.Articlelike{
		UserId: user.ID,
		ArticleId: input.ArticleId,
	}

	// いいねを登録
	err = repositorys.CreateLike(&newLike)

	if err != nil {
		return err
	}

	return nil

}

func DeleteLike(jwtToken string, articleId int) (error) {

	// 本人確認
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return err
	}

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(claims.Sub)
	if err != nil {
		return err
	}

	// いいねを削除
	err = repositorys.DeleteLike(user.ID, articleId)

	if err != nil {
		return err
	}

	return nil

}

func GetLikes(jwtToken string) ([]models.Article, error) {

	// 本人確認
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return nil, err
	}

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(claims.Sub)
	if err != nil {
		return nil, err
	}

	// いいねをした記事を取得
	articles, err := repositorys.GetLikes(user.ID)

	return articles, err

}

func GetLikesByArticleId(jwtToken string, articleId int) (bool, error) {

	// 本人確認
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return false, err
	}

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(claims.Sub)
	if err != nil {
		return false, err
	}

	// いいねをした記事を取得
	like, err := repositorys.GetLikesByArticleId(user.ID, articleId)

	if err != nil {
		return false, err
	}

	return like, nil

}