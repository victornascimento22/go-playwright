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
type MockVidaInfrastructure struct {
	mock.Mock
}

// ConnectDatabase mocks the database connection method.
// Returns a mocked *sql.DB instance and an error.
func (m *MockVidaInfrastructure) ConnectDatabase() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

// TestBuscarAniversariantesEmpresaHoje tests the BuscarAniversariantesEmpresa function to verify it retrieves employees with anniversaries correctly.
func TestBuscarAniversariantesVida(t *testing.T) {
	today := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}
	defer db.Close()

	mockInfrastructure := new(MockVidaInfrastructure)
	mockInfrastructure.On("ConnectDatabase").Return(db, nil)

	repository := repository.ConstructorAniversariantesVidaConnectionDatabase(mockInfrastructure)

	// Configurar o mock apenas com as colunas que realmente estão sendo selecionadas
	rows := sqlmock.NewRows([]string{
		"nome_cracha",
		"aniversario_vida",
		"url_aniversario_vida_tv",
	}).AddRow(
		"João",
		today,
		"http://example.com/joao",
	).AddRow(
		"Maria",
		today,
		"http://example.com/maria",
	)

	// Ajustar a query para corresponder exatamente à query real
	expectedQuery := `SELECT nome_cracha, TO_TIMESTAMP\(aniversario_vida, 'DD/MM/YYYY'\) as aniversario_vida, url_aniversario_vida_tv FROM DADOS_FUNCIONARIOS WHERE date_part\('day', to_date\(aniversario_vida, 'DD/MM/YYYY'\)\) = date_part\('day', CURRENT_DATE\) AND date_part\('month', to_date\(aniversario_vida, 'DD/MM/YYYY'\)\) = date_part\('month', CURRENT_DATE\)`

	mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

	aniversariantes, err := repository.GetAniversariantesVidaRepository()
	if err != nil {
		t.Fatalf("Erro ao buscar aniversariantes: %v", err)
	}

	// Verificações
	assert.NoError(t, err)
	assert.Len(t, aniversariantes, 2)

	// Verificar primeiro registro
	assert.Equal(t, "João", aniversariantes[0].NomeCracha)
	assert.Equal(t, today.Format("2006-01-02"), aniversariantes[0].AniversarioVida.Format("2006-01-02"))
	assert.Equal(t, "http://example.com/joao", aniversariantes[0].URLAniversarioVidaTv)

	// Verificar segundo registro
	assert.Equal(t, "Maria", aniversariantes[1].NomeCracha)
	assert.Equal(t, today.Format("2006-01-02"), aniversariantes[1].AniversarioVida.Format("2006-01-02"))
	assert.Equal(t, "http://example.com/maria", aniversariantes[1].URLAniversarioVidaTv)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
