package router

import (
	"md2s/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // ここで特定のオリジンを許可することもできます（例: []string{"http://localhost:3000"})
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッド
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 許可するヘッダー
		AllowCredentials: true, // クレデンシャルを許可するかどうか
		MaxAge:           12 * time.Hour, // キャッシュの最大時間
	}))



	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})


	users := r.Group("/users")
	users.GET("/:id", controllers.GetUsers)
	users.POST("", controllers.CreateUser)
	users.OPTIONS("", controllers.CreateUser)
	// users.DELETE("/:id", controllers.DeleteUser)

	imgs := r.Group("/imgs")
	imgs.POST("/imgs", controllers.UploadImg)

	// r.POST("/google", controllers.GoogleLogin)

	r.Run(":8080")
}