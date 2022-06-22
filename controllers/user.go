package controllers

import (
	"net/http"
	"time"

	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/httputil"
	"github.com/Andreasmalleus/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary      Show a user
// @Description  get user by id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  httputil.HttpError
// @Router       /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	err := config.Database.QueryRow(`SELECT * FROM "user" WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary      List users
// @Description  get users
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Failure      404  {object} httputil.HttpError
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	users := []models.User{}
	rows, err := config.Database.Query(`SELECT * FROM "user"`)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err.Error())
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
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary      Create a user
// @Description  create user with json
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body     models.User  true  "Create user"
// @Success      200  {object}  models.User
// @Failure      400  {object}  httputil.HttpError
// @Router       /user [post]
func CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	if user.Name == "" || user.Email == "" {
		httputil.NewError(c, http.StatusBadRequest, "Name and Email are required")
		return
	}
	currentTime := time.Now().Format(time.RFC3339)
	_, err := config.Database.Exec(`INSERT INTO "user" (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4)`, user.Name, user.Email, currentTime, currentTime)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  update user with json
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        user  body     models.UpdateUser  true  "Update user"
// @Success      200  {string}  models.User
// @Failure      400  {object}  httputil.HttpError
// @Router       /user/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := models.UpdateUser{}
	if err := c.ShouldBind(&body); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{}
	err := config.Database.QueryRow(`SELECT * FROM "user" WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}

	currentTime := time.Now().Format(time.RFC3339)

	if body.Email == "" {
		httputil.NewError(c, http.StatusBadRequest, "Email is required")
		return
	}
	_, exErr := config.Database.Exec(`UPDATE "user" SET email = $1, updated_at = $2 WHERE id = $3`, body.Email, currentTime, user.ID)
	if exErr != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  delete user with id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {string}  models.User
// @Failure      400  {object}  httputil.HttpError
// @Failure      404  {object}  httputil.HttpError
// @Router       /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res, err := config.Database.Exec(`DELETE FROM "user" WHERE id = $1`, id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	if rowsAffected == 0 {
		httputil.NewError(c, http.StatusNotFound, "No user found with that id")
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
