package dto

type GetArticlesData struct {
	// limitとoffsetはページネーションのためのパラメータ
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}