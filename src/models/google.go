package models

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Google公開鍵
type GoogleCerts struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
		Use string `json:"use"`
	} `json:"keys"`
}

// IDトークンのペイロード
type GoogleIDToken struct {
	Iss     string `json:"iss"`     // 発行者
	Aud     string `json:"aud"`     // 対象クライアントID
	Exp     int64  `json:"exp"`     // 有効期限
	Email   string `json:"email"`   // ユーザーのメールアドレス
	Name    string `json:"name"`    // ユーザー名
	Picture string `json:"picture"` // プロフィール画像 URL
	Sub     string `json:"sub"`     // ユーザーID
}

// GetAudience implements jwt.Claims.
func (g *GoogleIDToken) GetAudience() (jwt.ClaimStrings, error) {
	if g.Aud == "" {
		return nil, errors.New("audience is empty")
	}
	return jwt.ClaimStrings{g.Aud}, nil
}

// GetExpirationTime implements jwt.Claims.
func (g *GoogleIDToken) GetExpirationTime() (*jwt.NumericDate, error) {
	if g.Exp == 0 {
		return nil, errors.New("expiration time is not set")
	}
	return jwt.NewNumericDate(time.Unix(g.Exp, 0)), nil
}

// GetIssuedAt implements jwt.Claims.
func (g *GoogleIDToken) GetIssuedAt() (*jwt.NumericDate, error) {
	// IssuedAt (iat) が Google のトークンに存在する場合のみ実装
	// ここではダミー実装
	return nil, nil
}

// GetIssuer implements jwt.Claims.
func (g *GoogleIDToken) GetIssuer() (string, error) {
	if g.Iss == "" {
		return "", errors.New("issuer is empty")
	}
	return g.Iss, nil
}

// GetNotBefore implements jwt.Claims.
func (g *GoogleIDToken) GetNotBefore() (*jwt.NumericDate, error) {
	// NotBefore (nbf) が Google のトークンに存在する場合のみ実装
	// ここではダミー実装
	return nil, nil
}

// GetSubject implements jwt.Claims.
func (g *GoogleIDToken) GetSubject() (string, error) {
	if g.Sub == "" {
		return "", errors.New("subject is empty")
	}
	return g.Sub, nil
}
