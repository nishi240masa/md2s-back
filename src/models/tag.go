package models

type Tag struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Word string `json:"word"`
}