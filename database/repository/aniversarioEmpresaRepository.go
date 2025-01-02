// Package repository handles database operations and data access.
package repository

import (
	"fmt"

	infra "gitlab.com/applications2285147/api-go/infrastructure"
	"gitlab.com/applications2285147/api-go/internal/models"
)

// IAniversariantesEmpresaRepository defines an interface for fetching employees' anniversaries.
type IAniversariantesEmpresaRepository interface {
	// BuscarAniversariantesEmpresa retrieves a list of employees celebrating their work anniversary.
	BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error)
}

// IConnectDatabase encapsulates the database connection logic.
type IConnectDatabase struct {
	// infrastructure provides the methods to connect to the database.
	infrastructure infra.IConnectDatabase
}

// ConstructorConnectDatabase initializes a new instance of IConnectDatabase.
// i: an implementation of the IConnectDatabase interface.
// Returns a pointer to an IConnectDatabase instance.
func ConstructorConnectDatabase(i infra.IConnectDatabase) *IConnectDatabase {
	return &IConnectDatabase{
		infrastructure: i,
	}
}

// BuscarAniversariantesEmpresa retrieves a list of employees celebrating their work anniversary.
// It queries the database for employees whose work anniversary matches the current date.
// Returns a slice of Aniversariantes and an error, if any.
func (i *IConnectDatabase) BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error) {
	// SQL query to fetch employees celebrating their work anniversary today.
	query := `SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv
		FROM DADOS_FUNCIONARIOS
		WHERE date_part('day', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('day', CURRENT_DATE)
		AND date_part('month', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('month', CURRENT_DATE);`

	// Establish a connection to the database.
	db, err := i.infrastructure.ConnectDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Ensure the database connection is closed when the function exits.
	defer db.Close()

	// Execute the SQL query.
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	// Ensure the rows are closed when the function exits.
	defer rows.Close()

	// Initialize a slice to hold the result set.
	var aniversariantes []models.Aniversariantes

	// Iterate through the result set and populate the slice.
	for rows.Next() {
		var aniv models.Aniversariantes
		// Ensure the field names match the struct definition
		err := rows.Scan(&aniv.NomeCracha, &aniv.AniversarioEmpresa, &aniv.UrlAniversarioEmpresaTv)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		fmt.Printf("Nome: %s\nAniversario empresa: %s\n URL: %s\n", aniv.NomeCracha, aniv.AniversarioEmpresa, aniv.UrlAniversarioEmpresaTv)
		aniversariantes = append(aniversariantes, aniv)
	}

	// Return the list of employees and no error.
	return aniversariantes, nil
}
