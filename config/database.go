package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Andreasmalleus/go-rest-api/utils"
	"github.com/joho/godotenv"
)

var Database *sql.DB

func InitDatabase(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")

	connString := fmt.Sprintf("postgresql://%s@localhost:5432/%s?sslmode=disable", user, dbname)
	Database, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	err = Database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}

//populates the database with randomly generated mock data
func PopulateDatabase() {
	insertPostsSql := utils.GetFileContents("config/sql/insert-posts.sql")
	insertUsersSql := utils.GetFileContents("config/sql/insert-users.sql")
	_, usersErr := Database.Exec(insertUsersSql)
	if usersErr != nil {
		log.Fatal("PopulateDatabase error: ", usersErr)
	}
	_, postsErr := Database.Exec(insertPostsSql)
	if postsErr != nil {
		log.Fatal("PopulateDatabase error", postsErr)
	}
	log.Println("Database populated successfully")
}

func ClearDatabase() {
	_, err := Database.Exec(`TRUNCATE "user" CASCADE`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database cleared successfully")
}
