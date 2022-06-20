package controllers

import (
	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	err := config.Database.QueryRow(`SELECT * FROM "user" WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": user,
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
