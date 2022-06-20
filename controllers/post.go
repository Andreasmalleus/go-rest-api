package controllers

import (
	"time"

	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	err := config.Database.QueryRow(`SELECT * FROM post WHERE id = $1`, id).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"post": &post,
	})
}

func GetAllPosts(c *gin.Context) {
	posts := []models.Post{}
	rows, err := config.Database.Query(`SELECT * FROM post`)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	defer rows.Close()
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserId)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	c.JSON(200, gin.H{
		"posts": posts,
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
	id := c.Param("id")
	_, err := config.Database.Exec(`DELETE FROM post WHERE id = $1`, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "Post deleted successfully",
	})
}
