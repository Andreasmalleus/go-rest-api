package controllers

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "GET",
	})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": "GET",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "POST",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "PUT",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "DELETE",
	})
}
