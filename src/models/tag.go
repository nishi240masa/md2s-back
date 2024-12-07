package models

type Tag struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Word string `gorm:"column:word" json:"word"`
}
