package database

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"                // Certifique-se de importar o driver PostgreSQL
	"github.com/stretchr/testify/assert" // Importando a biblioteca de asserções
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
	"gitlab.com/applications2285147/api-go/database"
)

func TestBuscarAniversariantesDoDia(t *testing.T) {
	db, err := database.ConnectDatabase()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	aniversarioController := controller.ConstructorAniversarioController(db)

	aniversariantes, err := aniversarioController.BuscarAniversariantesDoDia()

	if err != nil {
		t.Fatalf("Failed to fetch aniversariantes: %v", err)
	}

	// Adicionando asserção para verificar se a lista de aniversariantes não está vazia
	assert.NotEmpty(t, aniversariantes, "A lista de aniversariantes não deve estar vazia")

	fmt.Println(aniversariantes)
}
