package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=muzzyaqow dbname=blog port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change the database provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Post{}) // register Post model

	DB = database
}

func main() {
	router := gin.Default()

	ConnectDatabase() // new!

	router.POST("/posts", CreatePost)
	router.GET("/posts", FindPosts)
	router.GET("/posts/:id", FindPost)
	router.PUT("/posts/:id", UpdatePost)
	router.DELETE("/posts/:id", DeletePost)

	router.Run("localhost:8080")
}
