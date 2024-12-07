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

type SearchArticlesData struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}