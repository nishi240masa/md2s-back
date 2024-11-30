package controllers

import (
	"md2s/dto"
	"md2s/repositorys"
	"md2s/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetUser(c *gin.Context){

	memberId,err := strconv.ParseUint(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		
	}

	result,err:=repositorys.GetUserById(uint(memberId))
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}



func CreateUser(c *gin.Context){
	var user dto.CreateUserData

	if error := c.ShouldBindJSON(&user); error != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":error.Error()})
		return
		
	}

	newUser,err:= services.CreateUser(user)


	if err != nil{

		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_pkey\" (SQLSTATE 23505)"{
			c.JSON(http.StatusBadRequest,gin.H{"error":"既に登録されているGoogleIDです"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
		
	}
	c.JSON(http.StatusCreated,newUser)
}

func DeleteUser(c *gin.Context){

	userId := c.Param("id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		
	}

	err = services.DeleteUser(uint(memberId))
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusNoContent,nil)
}

