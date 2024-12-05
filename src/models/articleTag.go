package models

type Articletagrelations struct {
	ID        int `gorm:"primary_key" json:"id"`
	ArticleId int `json:"article_id"`
	TagId     int `json:"tag_id"`
}