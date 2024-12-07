package models

import (
	"github.com/google/uuid"
)

type UUID = uuid.UUID

// StringToUUID
func StringToUUID(s string) (UUID, error) {
	return uuid.Parse(s)
}

type User struct {
	ID        UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	IconURL   string         `gorm:"type:varchar(255);not null" json:"icon_url"`
	GoogleId  string         `gorm:"type:varchar(50);unique;not null" json:"google_id"`
	QiitaId   string         `gorm:"type:varchar(50);unique" json:"qiita_id"`
}

type GetUser struct {
	ID      UUID   `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	Qiita_link bool `json:"qiita_link"`
	Total_get_like_count int `json:"total_get_like_count"`
	Total_posts_articles int `json:"total_article_count"`
}

type UserSortOptions struct {
	OrderBy string `json:"order_by"` // "name" or "created_at"
	Order   string `json:"order"`   // "asc" or "desc"
}
