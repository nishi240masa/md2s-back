package dto

type CreateTagData struct {
	Word string `gorm:"not null" json:"word"`
}