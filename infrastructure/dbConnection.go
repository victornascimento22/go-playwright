// Package infra provides infrastructure-related functionalities such as database connections.
package infra

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// IConnectDatabase defines an interface for connecting to a database.
type IConnectDatabase interface {
	// ConnectDatabase establishes a connection to the database and returns a database handle.
	ConnectDatabase() (*sql.DB, error)
}

// Database is a struct that implements the IConnectDatabase interface.
type Database struct{}

// ConstructorConnectDatabase creates a new instance of the Database struct.
// Returns an implementation of IConnectDatabase.
func ConstructorConnectDatabase() IConnectDatabase {
	return &Database{}
}

// ConnectDatabase loads environment variables, establishes a connection to the PostgreSQL database,
// and ensures the connection is active.
// Returns a pointer to the database handle or an error if the connection fails.
func (d *Database) ConnectDatabase() (*sql.DB, error) {
	// Load environment variables from the .env file.
	err := godotenv.Load("/home/loadt/api-go/.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Retrieve database credentials from environment variables.
	var (
		DbHost     = os.Getenv("DB_HOST")
		DbPort     = os.Getenv("DB_PORT")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASSWORD")
		DbName     = os.Getenv("DB_NAME")
	)

	// Create the connection string for PostgreSQL.
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)

	// Open a new connection to the PostgreSQL database.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// Verify the database connection.
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
