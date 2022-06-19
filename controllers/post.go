package controllers

import "github.com/gin-gonic/gin"

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
	c.JSON(200, gin.H{
		"post": "POST",
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
