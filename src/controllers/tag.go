package controllers

import (
	"md2s/dto"
	"md2s/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	tags, err := services.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
	var input dto.CreateTagData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := services.CreateTag(input)
	c.JSON(http.StatusOK, tag)
}