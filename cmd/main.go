package main

import (
	_ "github.com/lib/pq" // PostgreSQL driver
	api "gitlab.com/applications2285147/api-go/api/router"
	"gitlab.com/applications2285147/api-go/database"
)

func main() {
	db, err := database.ConnectDatabase() // ou o nome da sua função de conexão
	if err != nil {
		panic(err)
	}
	defer db.Close()
	api.Router(db)
}
