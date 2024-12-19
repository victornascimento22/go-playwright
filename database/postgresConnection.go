package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Valgard/godotenv"
)

var DB *sql.DB

func ConnectDatabase() (*sql.DB, error) {

	err := godotenv.Load("/home/victor/api-go-ssh/api-go/.env")

	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	var (
		DbHost     = os.Getenv("DB_HOST")
		DbPort     = os.Getenv("DB_PORT")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASSWORD")
		DbName     = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
