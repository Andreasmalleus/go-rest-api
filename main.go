package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Andreasmalleus/go-rest-api/config"
	"github.com/Andreasmalleus/go-rest-api/routes"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")

	connString := fmt.Sprintf("postgresql://%s@localhost:5432/%s?sslmode=disable", user, dbname)
	config.Database, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = config.Database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	router := routes.SetRouter()
	router.Run(":3000")
}
