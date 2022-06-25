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

func TestGetPost(t *testing.T) {
	const postId = 1
	req, _ := http.NewRequest("GET", URL+"/post/"+strconv.Itoa(postId), nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	var post models.Post
	json.Unmarshal(w.Body.Bytes(), &post)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, post)
}

func TestGetAllPosts(t *testing.T) {
	req, _ := http.NewRequest("GET", URL+"/posts", nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	var posts []models.Post
	json.Unmarshal(w.Body.Bytes(), &posts)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, posts)
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		Title:   "Test",
		Content: "Test",
		UserId:  3,
	}
	jsonData, _ := json.Marshal(&post)
	req, _ := http.NewRequest("POST", URL+"/post", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	createdPost := models.Post{}
	err := json.Unmarshal(w.Body.Bytes(), &createdPost)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, &createdPost)
	assert.Equal(t, &post, &createdPost)
}

func TestUpdatePost(t *testing.T) {
	const postId = 1
	updateBody := models.UpdatePost{
		Title:   "This is the test post",
		Content: "Test TEst TESt TEST",
	}
	jsonValue, _ := json.Marshal(&updateBody)
	req, _ := http.NewRequest("PUT", URL+"/post/"+strconv.Itoa(postId), bytes.NewReader(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePost(t *testing.T) {
	const postId = 1
	req, _ := http.NewRequest("DELETE", URL+"/post/"+strconv.Itoa(postId), nil)
	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
