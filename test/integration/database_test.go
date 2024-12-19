package database

import (
	"os"
	"testing"

	_ "github.com/lib/pq" // Certifique-se de importar o driver PostgreSQL
)

func TestConnectDatabase(t *testing.T) {
	// Configurar variáveis de ambiente de teste
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_NAME", "testdb")

	// Teste de conexão bem-sucedida
	t.Run("Successful connection", func(t *testing.T) {
		db, err := ConnectDatabase()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer db.Close()
	})

	// Teste de erro ao carregar o .env (remova ou renomeie o arquivo .env temporariamente para simular isso)
	t.Run("Missing .env file", func(t *testing.T) {
		originalPath := ".env"
		tempPath := ".env.bak"
		_ = os.Rename(originalPath, tempPath) // Renomeia o .env

		defer os.Rename(tempPath, originalPath) // Restaura o .env após o teste

		_, err := ConnectDatabase()
		if err == nil {
			t.Fatal("Expected error for missing .env file, got none")
		}
	})

	// Teste de erro ao conectar (com credenciais inválidas)
	t.Run("Invalid credentials", func(t *testing.T) {
		os.Setenv("DB_PASSWORD", "wrongpassword")
		_, err := ConnectDatabase()
		if err == nil {
			t.Fatal("Expected error for invalid credentials, got none")
		}
	})
}
