package models
type UUID [16]byte

type User struct {
	ID   UUID   `gorm:"primary_key auto_increment" json:"id"`
	Name string `json:"name"`
	IconURL string `json:"icon_url"`
	GoogleId string `json:"google_id"`
}

type UserCreate struct {
	Name string `json:"name"`
	IconURL string `json:"icon_url"`
	GoogleId string `json:"google_id"`
}

type like struct {
	id uint `gorm:"primary_key auto_increment" json:"id"`
	user_id UUID `json:"user_id"`
	article_id uint `json:"article_id"`
}