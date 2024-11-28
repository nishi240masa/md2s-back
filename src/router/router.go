package router

import (
	"md2s/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/member/:id", controllers.GetMember)

	r.Run(":8080")
}