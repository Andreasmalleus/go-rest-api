package controllers

import (
	"net/http"
	"time"

	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/httputil"
	"github.com/Andreasmalleus/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	err := config.Database.QueryRow(`SELECT * FROM post WHERE id = $1`, id).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.UserId)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {
	posts := []models.Post{}
	rows, err := config.Database.Query(`SELECT * FROM post`)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
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
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	post := models.CreatePost{}
	if err := c.ShouldBind(&post); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	_, err := config.Database.Exec(`INSERT INTO post (title, content, created_at, updated_at, user_id) VALUES ($1,$2,$3,$4,$5)`, &post.Title, &post.Content, currentTime, currentTime, &post.UserId)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	body := models.UpdatePost{}
	if err := c.ShouldBind(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	if body.Title == "" && body.Content == "" {
		httputil.NewError(c, http.StatusBadRequest, "No data to update")
		return
	}
	_, err := config.Database.Exec(`UPDATE post SET title = $1, content = $2, updated_at = $3 WHERE id = $4`, &body.Title, &body.Content, currentTime, id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	_, err := config.Database.Exec(`DELETE FROM post WHERE id = $1`, id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
