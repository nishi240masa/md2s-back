package models

import (
	"github.com/google/uuid"
)

type UUID = uuid.UUID

type User struct {
	ID        UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	IconURL   string         `gorm:"type:varchar(255);not null" json:"icon_url"`
	GoogleId  string         `gorm:"type:varchar(50);unique;not null" json:"google_id"`
	QiitaId   string         `gorm:"type:varchar(50);unique" json:"qiita_id"`
}

type UserSortOptions struct {
	OrderBy string `json:"order_by"` // "name" or "created_at"
	Order   string `json:"order"`   // "asc" or "desc"
}
