package models

type Articlelike struct {
	ID        int `gorm:"primary_key" json:"id"`
	ArticleId int `json:"article_id"`
	UserId    UUID `json:"user_id"`
}