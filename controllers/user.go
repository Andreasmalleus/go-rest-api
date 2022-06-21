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

type updateRequestBody struct {
	Email string `json:"email"`
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := updateRequestBody{}
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{}
	err := config.Database.QueryRow(`SELECT * FROM "user" WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	currentTime := time.Now().Format(time.RFC3339)

	if body.Email == "" {
		c.JSON(400, gin.H{
			"error": "Email is required",
		})
		return
	}
	if body.Email == "" {
		c.JSON(400, gin.H{
			"error": "Email is required",
		})
		return
	}
	_, exErr := config.Database.Exec(`UPDATE "user" SET email = $1, updated_at = $2 WHERE id = $3`, body.Email, currentTime, user.ID)
	if exErr != nil {
		c.JSON(400, gin.H{
			"error": exErr.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status": "User update successful...",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res, err := config.Database.Exec(`DELETE FROM "user" WHERE id = $1`, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		c.JSON(400, gin.H{
			"error": "No user found with that id",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "user successfuly deleted",
	})
}
