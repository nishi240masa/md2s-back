package controllers

import (
	"errors"
	"md2s/dto"
	"md2s/models"
	"md2s/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractJWTFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func GetUsers(c *gin.Context) {
	var sortOptions models.UserSortOptions
	if err := c.ShouldBindQuery(&sortOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sort options"})
		return
	}

	if sortOptions.OrderBy == "" {
		sortOptions.OrderBy = "created_at"
	}
	if sortOptions.Order == "" {
		sortOptions.Order = "desc"
	}

	users, err := services.GetUsers(sortOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserFromJWT(c *gin.Context) {
	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUserByJWT(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	jwtToken, err := extractJWTFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input dto.CreateUserData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.CreateUser(jwtToken, input)
	if err != nil {
		if err.Error() == "user already exists" {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
			
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
