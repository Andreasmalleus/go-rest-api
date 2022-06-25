package main

import (
	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/routes"
	_ "github.com/lib/pq"
)

// @title           RESTful API with GO
// @version         1.0
// @description     RESTful API created with GO for learning purposes
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	config.InitDatabase(".env")
	router := routes.SetRouter()
	router.Run(":3000")
}
