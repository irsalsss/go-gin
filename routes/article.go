package routes

import (
	"go-gin-rest/config"
	"go-gin-rest/models"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item models.Article

	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}
}

func PostArticle(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil ngepost",
		"data":   "item",
	})
}
