package services

import (
	"errors"
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
	"md2s/utils/qitta"
	"net/http"
)

func AlignmentQiita(jwtToken string, input dto.AlignmentQiita) (*models.User, error) {

	// JWTトークンからユーザー情報を取得
    claims, err := VerifyGoogleToken(jwtToken)
    if err != nil {
        return nil, err
    }

	userId := claims.Sub

	// 既に登録されているか確認
	user, err := repositorys.GetUserByGoogleID(userId)
	if err != nil {
		return nil, err
	}

	// Qiitaのアクセストークンを取得
		token, err := qitta.GetQiitaAccessToken(input.QitaCode) // 正しいフィールド名を使用
		if err != nil {
			return nil, err
		}

	user.QiitaId = token

	// ユーザー情報を更新
	err = repositorys.AlignmentQiita(user)
	if err != nil {
		return nil, err
	}

	return user, nil



}

func GetQiitaArticles(jwtToken string) (*http.Request, error) {
	// JWTトークンからユーザー情報を取得
	claims, err := VerifyGoogleToken(jwtToken)
	if err != nil {
		return nil, err
	}

	userId := claims.Sub

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(userId)
	if err != nil {
		return nil, err
	}

	if user.QiitaId == user.GoogleId {
		return nil, errors.New("Qiitaアカウントが連携されていません")
	}


	// Qiitaのアクセストークンを取得
	token := user.QiitaId

	// Qiitaの記事を取得
	articles, err := qitta.GetQiitaArticles(token)
	if err != nil {
		return nil, err
	}

	return articles, nil
}