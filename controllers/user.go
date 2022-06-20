package controllers

import (
	"time"

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
	users := []models.User{}
	rows, err := config.Database.Query(`SELECT * FROM "user"`)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Name == "" || user.Email == "" {
		c.JSON(400, gin.H{
			"error": "Name and Email are required",
		})
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	_, err := config.Database.Exec(`INSERT INTO "user" (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4)`, user.Name, user.Email, currentTime, currentTime)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
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
