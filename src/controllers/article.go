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

