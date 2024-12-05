package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
	"md2s/utils/qitta"
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