package controllers

import (
	"md2s/dto"
	"md2s/services"
	"net/http"
	"strconv"

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

func GetTag(c *gin.Context) {

	id,_ := strconv.Atoi(c.Param("id"))
	tag, err := services.GetTag(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	var input []dto.CreateTagData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag,err := services.CreateTag(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, tag)
}

func UpdateTag(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	var input dto.CreateTagData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := services.UpdateTag(id, input)

	if tag != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tag.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteTag(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	tag := services.DeleteTag(id)

	if tag != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tag.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}