package config

import (
	"fmt"
	"os"

	"github.com/irsalsss/go-gin/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.User{}).Related(&models.Article{})
}
