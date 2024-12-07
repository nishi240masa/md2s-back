package controllers

import (
	"md2s/dto"
	"md2s/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSlide(c *gin.Context) {

	var input  dto.RequestBody

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	slide, err := services.GetSlide(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"slide": slide})
}
