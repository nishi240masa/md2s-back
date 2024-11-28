package models

type Member struct {
	ID   uint   `gorm:"primary_key auto_increment" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}