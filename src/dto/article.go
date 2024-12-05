package dto

type GetArticlesData struct {
	// limitとoffsetはページネーションのためのパラメータ
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type CreateArticleData struct {
	Title        string `json:"title"`
	MainMD       string `json:"main_MD"`
	SlideMD      string `json:"slide_MD"`
	Public       bool   `json:"public"`
	QiitaArticle bool   `json:"qiita_article"`
	Tags         []struct {
		ID   int `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`

}

// [
//   "create_user_id":"387a4e2d-5b8d-4d-9035-ee95680b66b4",
//   "title": "title1",
//   "main_MD":"本文のMarkdown",
//   "slide_MD":"スライドのMarkdown",
//   "public": true
//   "qiita_article": false,
//   "tags": [
//     {
//       "id": "1",
//       "name": "tag1"
//     },
//     {
//       "id": "2",
//       "name": "tag2"
//     }
//   ]
// ]