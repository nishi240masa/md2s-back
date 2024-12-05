package models

import "time"



type Article struct {
    ID          int       `gorm:"primary_key" json:"id"`
    UserId      UUID    `db:"user_id"`
    Title       string    `db:"title"`
    MainMD      string    `db:"main_md"`
    SlideMD     *string   `db:"slide_md"`
    CreatedAt   time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
    LikeCount   int       `db:"like_count"`
    Public      bool      `db:"public"`
    QiitaArticle bool     `db:"qiita_article"`
}

