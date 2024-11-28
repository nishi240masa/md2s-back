package controllers

import (
	"md2s/repositorys"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetMember(c *gin.Context){

	memberId,err := strconv.ParseUint(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		
	}

	result,err:=repositorys.GetMemberById(uint(memberId))
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}
