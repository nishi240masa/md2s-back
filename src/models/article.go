package models

import "time"

type Article struct {
    ID          int       `gorm:"primary_key" json:"id"`
    UserId      UUID    `db:"user_id" gorm:"column:user_id" json:"user_id"`
	Name  string    `db:"name" gorm:"column:name" json:"name"`
	IconURL string    `db:"icon_url" json:"icon_url"`
    Title       string    `db:"title" gorm:"column:title" json:"title"`
    MainMD      string    `db:"main_md" gorm:"column:main_md" json:"main_MD"`
    SlideMD     *string   `db:"slide_md" gorm:"column:slide_md" json:"slide_MD"`
    CreatedAt   time.Time `db:"created_at" gorm:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" gorm:"updated_at" json:"updated_at"`
    LikeCount   int       `db:"like_count" gorm:"column:like_count" json:"like_count"`
    Public      bool      `db:"public" gorm:"column:public" json:"public"`
    QiitaArticle bool     `db:"qiita_article" gorm:"column:qiita_article" json:"qiita_article"`
	Tags		[]Tag     `gorm:"many2many:articletagrelations;" json:"tags"`

}

type CreateArticle struct {
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

type Articles struct {
	ID          int       `gorm:"primary_key" json:"id"`
    UserId      UUID    `db:"user_id" gorm:"column:user_id" json:"user_id"`
    Title       string    `db:"title" gorm:"column:title" json:"title"`
    MainMD      string    `db:"main_md" gorm:"column:main_md" json:"main_md"`
    SlideMD     *string   `db:"slide_md" gorm:"column:slide_md" json:"slide_md"`
    CreatedAt   time.Time `db:"created_at" gorm:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" gorm:"updated_at" json:"updated_at"`
    LikeCount   int       `db:"like_count" gorm:"column:like_count" json:"like_count"`
    Public      bool      `db:"public" gorm:"column:public" json:"public"`
    QiitaArticle bool     `db:"qiita_article" gorm:"column:qiita_article" json:"qiita_article"`
}