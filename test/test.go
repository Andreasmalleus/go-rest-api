package test

import (
	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitDatabase("../.env")
	Router = routes.SetRouter()
}

const URL = "/api/v1"

var Router *gin.Engine
