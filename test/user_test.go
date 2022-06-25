package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Andreasmalleus/go-rest-api/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	const userId = 1
	req, _ := http.NewRequest("GET", URL+"/user/"+strconv.Itoa(userId), nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	user := models.User{}
	json.Unmarshal(w.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
}

func TestGetAllUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", URL+"/users", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	users := []models.User{}
	json.Unmarshal(w.Body.Bytes(), &users)
	log.Println("users", &users)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}

func TestCreateUser(t *testing.T) {
	user := models.User{
		Name:  "TestUser",
		Email: "TestUser@hotmail.com",
	}
	jsonData, _ := json.Marshal(&user)
	req, _ := http.NewRequest("POST", URL+"/user", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	createdUser := models.User{}
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, &createdUser)
	assert.Equal(t, &user, &createdUser)
}

func TestUpdateUser(t *testing.T) {
	const userId = 1
	updateBody := models.UpdateUser{
		Email: "Test@hotmail.com",
	}
	jsonValue, _ := json.Marshal(&updateBody)
	req, _ := http.NewRequest("PUT", URL+"/user/"+strconv.Itoa(userId), bytes.NewReader(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	const userId = 1
	req, _ := http.NewRequest("DELETE", URL+"/user/"+strconv.Itoa(userId), nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
