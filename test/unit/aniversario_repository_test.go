// controller/aniversario_controller_test.go
package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/applications2285147/api-go/database/repository"
	// Alteração para 'infra'
)

type MockInfrastructure struct {
	mock.Mock
}

func (m *MockInfrastructure) ConnectDatabase() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

func TestBuscarAniversariantesEmpresaHoje(t *testing.T) {
	// Pegando a data de hoje
	today := time.Now().Format("02/01/2006") // Formato: DD/MM/YYYY

	// Criando o mock do banco de dados
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do banco de dados: %v", err)
	}

	// Mock da conexão com o banco
	mockInfrastructure := new(MockInfrastructure)
	mockInfrastructure.On("ConnectDatabase").Return(db, nil)

	// Criando o repositório com a infraestrutura mockada
	repository := repository.ConstructorConnectDatabase(mockInfrastructure) // Passando a infraestrutura mockada

	// Definindo o comportamento esperado da consulta SQL, incluindo a data de hoje
	rows := sqlmock.NewRows([]string{"nome_cracha", "aniversario_empresa", "url_aniversario_empresa_tv"}).
		AddRow("João", today, "http://example.com/joao").
		AddRow("Maria", today, "http://example.com/maria")

	// Definindo que o método db.Query será chamado e retornará as linhas simuladas
	mock.ExpectQuery(`SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv FROM DADOS_FUNCIONARIOS`).
		WillReturnRows(rows)

	// Chamando o método a ser testado
	aniversariantes, err := repository.BuscarAniversariantesEmpresa()

	// Verificando se ocorreu algum erro inesperado
	assert.NoError(t, err)

	// Verificando os resultados
	assert.Len(t, aniversariantes, 2)
	assert.Equal(t, "João", aniversariantes[0].Nome_cracha)
	assert.Equal(t, today, aniversariantes[0].Aniversario_empresa)
	assert.Equal(t, "http://example.com/joao", aniversariantes[0].Url_aniversario_empresa_tv)

	assert.Equal(t, "Maria", aniversariantes[1].Nome_cracha)
	assert.Equal(t, today, aniversariantes[1].Aniversario_empresa)
	assert.Equal(t, "http://example.com/maria", aniversariantes[1].Url_aniversario_empresa_tv)

	// Verificando se todas as expectativas do mock foram atendidas
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
