package controllers

import (
	"md2s/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSlide(c *gin.Context) {
	var requestBody struct {
		Input string `json:"md"` // リクエストボディのJSONフィールド
	}

	// JSONのバインド
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	slide, err := services.GetSlide([]byte(requestBody.Input))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"slide": slide})
}
