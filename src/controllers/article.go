package controllers

import (
	"md2s/dto"
	"md2s/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {

	limit,_ := strconv.Atoi(c.Query("limit"))
	offset,_ := strconv.Atoi(c.Query("offset"))



	input := dto.GetArticlesData{
		Limit: limit,
		Offset: offset,
	}

	articles, err := services.GetArticles(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)

}

func GetArticle(c *gin.Context) {
	
	id,_ := strconv.Atoi(c.Param("id"))

	article, err := services.GetArticle(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func CreateArticle(c *gin.Context) {
	var input dto.CreateArticleData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 本人確認
	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims, err := services.VerifyGoogleToken(jwtToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	googleId := claims.Sub



	// 記事作成
	article := services.CreateArticle(input, googleId)

	if article != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": article.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}


// update

func UpdateArticle(c *gin.Context) {

	id,_ := strconv.Atoi(c.Param("id"))


	var input dto.CreateArticleData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 本人確認
	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims, err := services.VerifyGoogleToken(jwtToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	googleId := claims.Sub

	// 記事更新
	article := services.UpdateArticle(id, input, googleId)

	if article != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": article.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
