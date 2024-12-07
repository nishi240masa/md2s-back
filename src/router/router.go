package router

import (
	"fmt"
	"md2s/controllers"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	// 環境変数 PORT を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // ここで特定のオリジンを許可することもできます（例: []string{"http://localhost:3000"})
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッド
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 許可するヘッダー
		AllowCredentials: true,                                                // クレデンシャルを許可するかどうか
		MaxAge:           12 * time.Hour,                                      // キャッシュの最大時間
	}))

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "connection success",
		})

	})

	// ユーザー
	users := r.Group("/users")
	users.GET("/:id", controllers.GetUsers)
	users.POST("", controllers.CreateUser)
	users.OPTIONS("", controllers.CreateUser)
	// users.DELETE("/:id", controllers.DeleteUser)

	// 画像
	imgs := r.Group("/imgs")
	imgs.POST("/imgs", controllers.UploadImg)

	// Qiita
	qiita := r.Group("/qiita")
	qiita.POST("", controllers.AlignmentQiita)

	// 記事
	articles := r.Group("/articles")
	articles.GET("", controllers.GetArticles)
	articles.GET("/:id", controllers.GetArticle)

	articles.POST("", controllers.CreateArticle)
	articles.PUT("/:id", controllers.UpdateArticle)

	// タグ
	tags := r.Group("/tags")
	tags.GET("", controllers.GetTags)
	tags.GET("/:id", controllers.GetTag)
	tags.POST("", controllers.CreateTag)
	tags.PUT("/:id", controllers.UpdateTag)
	tags.DELETE("/:id", controllers.DeleteTag)

	// like
	likes := r.Group("/likes")
	likes.GET("", controllers.GetLikes)
	likes.GET("/:id", controllers.GetLikesByArticleId)

	likes.POST("/", controllers.CreateLike)
	likes.DELETE("/:id", controllers.DeleteLike)

	// slide
	slide := r.Group("/slide")
	slide.GET("", controllers.GetSlide)

	// 指定されたポートでサーバーを開始
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
