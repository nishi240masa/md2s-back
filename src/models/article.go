package models

import "time"

// {
//     "id": "1",
//     "create_user_id": "387a4e2d-5b8d-4d-9035-ee95680b66b4",
//     "user_name": "user1",
//     "user_icon": "https://example.com/icon.jpg",
//     "title": "title1",
//     "main_MD": "本文のMarkdown",
//     "slide_MD": "スライドのMarkdown",
//     "created_at": "2021-01-01T00:00:00Z",
//     "updated_at": "2021-01-01T00:00:00Z",
//     "like_count": 10,
//     "public": true,
//     "qiita_article": true,
//     "tags": [
//       {
//         "id": "1",
//         "name": "tag1"
//       },
//       {
//         "id": "2",
//         "name": "tag2"
//       }
//     ]
//   },



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
	Tags		[]Tag     `gorm:"many2many:articletagrelations;"` 

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


