package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var posts = make(map[uint64]Post)

func main() {
	router := gin.Default()

	router.GET("/get/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		if _, ok := posts[id]; !ok {
			c.JSON(404, gin.H{"error": "Post not found"})
			return
		} else {
			c.JSON(200, posts[id])
		}
	})

	router.POST("/post", func(c *gin.Context) {
		var input Post
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		post := Post{
			ID:        uint64(len(posts) + 1),
			Title:     input.Title,
			Content:   input.Content,
			CreatedAt: time.Now(),
		}
		posts[post.ID] = post
		c.JSON(http.StatusOK, gin.H{"data": post})
	})

	router.PUT("/put/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		if _, ok := posts[id]; !ok {
			c.JSON(404, gin.H{"error": "Post not found"})
			return
		} else {
			var input Post
			if err := c.ShouldBindJSON(&input); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			posts[id] = Post{
				ID:        id,
				Title:     input.Title,
				Content:   input.Content,
				CreatedAt: posts[id].CreatedAt,
			}
			c.JSON(200, gin.H{"data": posts[id]})
		}
	})

	router.DELETE("/delete/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		if _, ok := posts[id]; !ok {
			c.JSON(404, gin.H{"error": "Post not found"})
			return
		} else {
			delete(posts, id)
			c.JSON(200, "Post deleted")
		}
	})

	router.Run("localhost:8080")
}
