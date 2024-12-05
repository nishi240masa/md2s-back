package models

type Article struct {
	ID          int    `json:"id"`
	UserId     int    `json:"user_id"`
	Title	   string `json:"title"`
	MainMD	   string `json:"main_md"`
	SlideMd	   string `json:"slide_md"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Like_count int    `json:"like_count"`
	Public bool   `json:"published"`
	QiitaId string `json:"qiita_id"`
}

