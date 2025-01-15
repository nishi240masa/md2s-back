package controllers

import (
	"fmt"
	"md2s/dto"
	"md2s/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlignmentQiita(c *gin.Context) {

	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input dto.AlignmentQiita
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.AlignmentQiita(jwtToken, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetQiitaArticles(c *gin.Context) {
	
	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles, err := services.GetQiitaArticles(jwtToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(articles)

	c.JSON(http.StatusOK, articles)

}