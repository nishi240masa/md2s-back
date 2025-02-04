package services

import (
	"md2s/dto"
	"md2s/models"
	"md2s/repositorys"
)


func GetArticles( quary dto.GetArticlesData) ([]models.Article, error) {

	

	articles, err := repositorys.GetArticles(quary)
	if err != nil {
		return nil, err
	}

	// 記事IDからタグIDを取得
	for i, article := range articles {
	tags, err := repositorys.GetArticleTagByArticleID(article.ID)
	if err != nil {
		return nil, err
	}


	var articleTags []models.Tag
	// タグIDからタグ情報を取得
	for _, tagRelation := range tags {
		tag, err := repositorys.GetTag(tagRelation.TagId)
		if err != nil {
			return nil, err
		}
		articleTags = append(articleTags, *tag)
	}
	article.Tags = articleTags
	articles[i] = article
	}


	return articles, nil
}

// 記事IDから記事情報を取得
func GetArticle(id int) (*models.Article, error) {
	article,err := repositorys.GetArticle(id)

	if err != nil {
		return nil, err
	}

	// 記事IDからタグIDを取得
	tags, err := repositorys.GetArticleTagByArticleID(article.ID)
	if err != nil {
		return nil, err
	}

	var articleTags []models.Tag
	// タグIDからタグ情報を取得
	for _, tagRelation := range tags {
		tag, err := repositorys.GetTag(tagRelation.TagId)
		if err != nil {
			return nil, err
		}
		articleTags = append(articleTags, *tag)
	}
	article.Tags = articleTags

	return article, nil
}

// ユーザーIDから記事情報を取得
func GetArticlesByUserId(user_id models.UUID) ([]models.Article, error) {
	
	articles, err := repositorys.GetArticlesByUserId(user_id)
	if err != nil {
		return nil, err
	}


		// 記事IDからタグIDを取得
		for i, article := range articles {
			tags, err := repositorys.GetArticleTagByArticleID(article.ID)
			if err != nil {
				return nil, err
			}
		
		
			var articleTags []models.Tag
			// タグIDからタグ情報を取得
			for _, tagRelation := range tags {
				tag, err := repositorys.GetTag(tagRelation.TagId)
				if err != nil {
					return nil, err
				}
				articleTags = append(articleTags, *tag)
			}
			article.Tags = articleTags
			articles[i] = article
			}
	return articles, nil
}

func CreateArticle(input dto.CreateArticleData,googleId string) error {

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(googleId)
	if err != nil {
		return err
	}
	
	newArticle := models.Articles{
		UserId: user.ID,
		Title: input.Title,
		MainMD: input.MainMD,
		SlideMD: &input.SlideMD,
		LikeCount: 0,
		Public: input.Public,
		QiitaArticle: input.QiitaArticle,
	}


	newId , err := repositorys.CreateArticle(&newArticle)

	if err != nil {
		return err
	}

	// タグの登録
	for _, tag := range input.Tags {

		article_id := newId
		tag_id := tag.ID

		articleTag := &models.Articletagrelations{
			ArticleId: article_id,
			TagId:     tag_id,
		}
		err := repositorys.CreateArticleTag(articleTag)
		if err != nil {
			return err
		}
	}

	return nil
}

func SearchArticles(input dto.SearchArticlesData) ([]models.Article, error) {

	// keywordを含む記事情報を取得
	articles, err := repositorys.SearchArticles(input)
	if err != nil {
		return nil, err
	}

	for i, article := range articles {
		tags, err := repositorys.GetArticleTagByArticleID(article.ID)
		if err != nil {
			return nil, err
		}
	
	
		var articleTags []models.Tag
		// タグIDからタグ情報を取得
		for _, tagRelation := range tags {
			tag, err := repositorys.GetTag(tagRelation.TagId)
			if err != nil {
				return nil, err
			}
			articleTags = append(articleTags, *tag)
		}
		article.Tags = articleTags
		articles[i] = article
		}

	return articles, nil

}

func UpdateArticle(id int, input dto.CreateArticleData,googleId string ) error {

	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(googleId)
	if err != nil {
		return err
	}

	article, err := repositorys.GetArticle(id)
	if err != nil {
		return err
	}

	// 記事の取得
	nowArticle, err := GetArticle(id)
	if err != nil {
		return err
	}

	// ユーザー確認
	if nowArticle.UserId != user.ID {
		return nil
	}

	// 記事の更新
	newArticle := models.Articles{
		ID: id,
		UserId: user.ID,
		Title: input.Title,
		MainMD: input.MainMD,
		SlideMD: &input.SlideMD,
		CreatedAt: nowArticle.CreatedAt,
		LikeCount: article.LikeCount,
		Public: input.Public,
		QiitaArticle: input.QiitaArticle,
	}

	err = repositorys.UpdateArticle(&newArticle)
	if err != nil {
		return err
	}


	// タグの更新
	// タグの削除
	err = repositorys.DeleteArticleTagByArticleID(id)
	if err != nil {
		return err
	}

	// タグの登録
	for _, tag := range input.Tags {
		article_id := id
		tag_id := tag.ID

		articleTag := &models.Articletagrelations{
			ArticleId: article_id,
			TagId:     tag_id,
		}
		err := repositorys.CreateArticleTag(articleTag)
		if err != nil {
			return err
		}
	}

	return nil


}


func DeleteArticle(id int,googleId string) error {
	
	// ユーザー情報を取得
	user, err := repositorys.GetUserByGoogleID(googleId)
	if err != nil {
		return err
	}

	// 記事の取得
	nowArticle, err := GetArticle(id)
	if err != nil {
		return err
	}

	// ユーザー確認
	if nowArticle.UserId != user.ID {
		return nil
	}

	// 記事の削除
	err = repositorys.DeleteArticle(id)

	if err != nil {
		return err
	}

	// タグの削除
	err = repositorys.DeleteArticleTagByArticleID(id)

	if err != nil {
		return err
	}

	return nil
}