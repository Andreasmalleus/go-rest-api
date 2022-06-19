package routes

import (
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return router
}
