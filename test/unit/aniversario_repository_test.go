// Package repository contains the test cases for the aniversario repository functionality.
package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/applications2285147/api-go/database/repository"
)

// MockInfrastructure is a mocked implementation of the infrastructure interface used to connect to the database.
type MockInfrastructure struct {
	mock.Mock
}

// ConnectDatabase mocks the database connection method.
// Returns a mocked *sql.DB instance and an error.
func (m *MockInfrastructure) ConnectDatabase() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

// TestBuscarAniversariantesEmpresaHoje tests the BuscarAniversariantesEmpresa function to verify it retrieves employees with anniversaries correctly.
func TestBuscarAniversariantesEmpresaHoje(t *testing.T) {
	// Fetch the current date in the format DD/MM/YYYY to simulate today's anniversaries.
	today := time.Now().Format("02/01/2006")

	// Create a mock database connection using sqlmock.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err) // Fail the test if the mock creation fails.
	}

	// Create a mock infrastructure for the database connection.
	mockInfrastructure := new(MockInfrastructure)
	mockInfrastructure.On("ConnectDatabase").Return(db, nil) // Define behavior for ConnectDatabase.

	// Initialize the repository with the mocked infrastructure.
	repository := repository.ConstructorConnectDatabase(mockInfrastructure)

	// Define the expected SQL query behavior, including mock data for today's anniversaries.
	rows := sqlmock.NewRows([]string{"nome_cracha", "aniversario_empresa", "url_aniversario_empresa_tv"}).
		AddRow("João", today, "http://example.com/joao").
		AddRow("Maria", today, "http://example.com/maria")

	// Expect the specific query to be called and return the mock rows.
	mock.ExpectQuery(`SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv FROM DADOS_FUNCIONARIOS`).
		WillReturnRows(rows)

	// Call the method being tested.
	aniversariantes, err := repository.BuscarAniversariantesEmpresa()

	// Assert no unexpected errors occurred.
	assert.NoError(t, err)

	// Assert the correct number of results were returned.
	assert.Len(t, aniversariantes, 2)

	// Assert the first returned record matches the expected mock data.
	assert.Equal(t, "João", aniversariantes[0].Nome_cracha)
	assert.Equal(t, today, aniversariantes[0].Aniversario_empresa)
	assert.Equal(t, "http://example.com/joao", aniversariantes[0].Url_aniversario_empresa_tv)

	// Assert the second returned record matches the expected mock data.
	assert.Equal(t, "Maria", aniversariantes[1].Nome_cracha)
	assert.Equal(t, today, aniversariantes[1].Aniversario_empresa)
	assert.Equal(t, "http://example.com/maria", aniversariantes[1].Url_aniversario_empresa_tv)

	// Verify that all expectations for the mock were met.
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err) // Assert that all defined expectations for the mock database were fulfilled.
}
