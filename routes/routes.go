package routes

import (
	"github.com/Andreasmalleus/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/user/:id", controllers.GetUser)
		v1.POST("/user", controllers.CreateUser)
		v1.PUT("/user/:id", controllers.UpdateUser)
		v1.DELETE("/user/:id", controllers.DeleteUser)

		v1.GET("/posts", controllers.GetAllPosts)
		v1.GET("/post/:id", controllers.GetPost)
		v1.POST("/post", controllers.CreatePost)
		v1.PUT("/post/:id", controllers.UpdatePost)
		v1.DELETE("/post/:id", controllers.DeletePost)
	}
	return router
}
