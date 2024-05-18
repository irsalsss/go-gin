package main

import (
	"go-gin-rest/config"
	"go-gin-rest/middleware"
	"go-gin-rest/routes"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)

		v1.GET("/profile", middleware.IsAuth(), routes.GetProfile)

		v1.GET("/article/:slug", routes.GetArticle)

		articles := v1.Group("/articles")
		{
			articles.GET("/", routes.GetHome)

			articles.GET("/tag/:tag", routes.GetArticleByTag)
			articles.POST("/", middleware.IsAuth(), routes.PostArticle)
			articles.PUT("/update/:id", middleware.IsAuth(), routes.UpdateArticle)
			articles.DELETE("/delete/:id", middleware.IsAdmin(), routes.DeleteArticle)
		}
	}

	router.Run()
}
