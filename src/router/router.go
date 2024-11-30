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


	users := r.Group("/users")
	r.GET("/user/:id", controllers.GetUser)
	

	r.POST("/users", controllers.CreateUser)

	r.POST("/imgs", controllers.UploadImg)

	r.Run(":8080")
}