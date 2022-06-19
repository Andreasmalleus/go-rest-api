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
