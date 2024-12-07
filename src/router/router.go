package router

import (
	"fmt"
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

		// クエリパラメータを取得
		code := c.Query("code")

		fmt.Println(code)

		// リダイレクト先のURLを生成
		redirectURL := "https://md2s-test-aa.vercel.app/"

		// リダイレクト
		c.Redirect(302, redirectURL)

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
	articles.GET("/search", controllers.SearchArticles)

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


	r.Run()
}