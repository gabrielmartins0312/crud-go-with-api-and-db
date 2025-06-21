package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	godotenv.Load()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error testing connection to the bank: ", err)
	}

	log.Println("Database connection successful!")
}
