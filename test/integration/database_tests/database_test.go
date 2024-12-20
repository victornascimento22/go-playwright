package database

import (
	"testing"

	_ "github.com/lib/pq" // Certifique-se de importar o driver PostgreSQL
	"gitlab.com/applications2285147/api-go/database"
)

func TestConnectDatabase(t *testing.T) {
	t.Run("Successful connection", func(t *testing.T) {
		db, err := database.ConnectDatabase()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer db.Close()
	})
}
