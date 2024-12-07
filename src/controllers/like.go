package controllers

import (
	"md2s/dto"
	"md2s/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateLike(c *gin.Context) {

	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input dto.Like
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateLike(jwtToken, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func DeleteLike (c *gin.Context) {

	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	err = services.DeleteLike(jwtToken, articleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

func GetLikes (c *gin.Context) {

	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles, err := services.GetLikes(jwtToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)

}

func GetLikesByArticleId (c *gin.Context) {

	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	like, err := services.GetLikesByArticleId(jwtToken, articleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, like)

}