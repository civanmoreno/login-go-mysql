package utilities

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error verifying connection to database: ", err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
