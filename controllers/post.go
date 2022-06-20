package controllers

import (
	"time"

	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"post": "GET",
	})
}

func GetAllPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": "GET",
	})
}

func CreatePost(c *gin.Context) {
	post := models.Post{}
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	_, err := config.Database.Exec(`INSERT INTO post (title, content, created_at, updated_at, user_id) VALUES ($1,$2,$3,$4,$5)`, &post.Title, &post.Content, currentTime, currentTime, &post.UserId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Post created successfully",
	})
}

func UpdatePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"post": "PUT",
	})
}

func DeletePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"post": "DELETE",
	})
}
