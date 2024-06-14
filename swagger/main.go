package main

import (
	"net/http"
	"strconv"
	"time"

	_ "example/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample CRUD server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/

type Post struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var posts = make(map[uint64]Post)

// @Summary Get a post by ID
// @Description Get a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint64 true "Post ID"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /get/{id} [get]
func getPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	post, ok := posts[id]
	if !ok {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary Create a new post
// @Description Create a new post with title and content
// @Tags posts
// @Accept json
// @Produce json
// @Param post body Post true "Create Post"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse
// @Router /post [post]
func createPost(c *gin.Context) {
	var input Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	post := Post{
		ID:        uint64(len(posts) + 1),
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}
	posts[post.ID] = post
	c.JSON(http.StatusOK, post)
}

// @Summary Update a post by ID
// @Description Update a post's title and content by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint64 true "Post ID"
// @Param post body Post true "Update Post"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /put/{id} [put]
func updatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	post, ok := posts[id]
	if !ok {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Post not found"})
		return
	}

	var input Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	posts[id] = post

	c.JSON(http.StatusOK, post)
}

// @Summary Delete a post by ID
// @Description Delete a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path uint64 true "Post ID"
// @Success 200 {string} string "Post deleted"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /delete/{id} [delete]
func deletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID"})
		return
	}

	if _, ok := posts[id]; !ok {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Post not found"})
		return
	}

	delete(posts, id)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/get/:id", getPost)
	router.POST("/post", createPost)
	router.PUT("/put/:id", updatePost)
	router.DELETE("/delete/:id", deletePost)

	router.Run("localhost:8080")
}
