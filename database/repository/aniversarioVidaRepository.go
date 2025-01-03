package repository

import (
	"fmt"

	infra "gitlab.com/applications2285147/api-go/infrastructure"
	"gitlab.com/applications2285147/api-go/internal/models"
)

type IAniversariantesVidaRepository interface {
	GetAniversariantesVidaRepository() ([]models.Aniversariantes, error)
}

type IAniversariantesVidaConnectionDatabase struct {
	// infrastructure provides the methods to connect to the database.
	infrastructure infra.IConnectDatabase
}

// ConstructorConnectDatabase initializes a new instance of IConnectDatabase.
// i: an implementation of the IConnectDatabase interface.
// Returns a pointer to an IConnectDatabase instance.
func ConstructorAniversariantesVidaConnectionDatabase(i infra.IConnectDatabase) *IAniversariantesVidaConnectionDatabase {
	return &IAniversariantesVidaConnectionDatabase{
		infrastructure: i,
	}
}

func (v *IAniversariantesVidaConnectionDatabase) GetAniversariantesVidaRepository() ([]models.Aniversariantes, error) {

	// SQL query to fetch employees celebrating their work anniversary today.
	query := `SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv
		FROM DADOS_FUNCIONARIOS
		WHERE date_part('day', to_date(aniversario_vida, 'DD/MM/YYYY')) = date_part('day', CURRENT_DATE)
		AND date_part('month', to_date(aniversario_vida, 'DD/MM/YYYY')) = date_part('month', CURRENT_DATE);`

	// Establish a connection to the database.
	db, err := v.infrastructure.ConnectDatabase()
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
		err := rows.Scan(&aniv.NomeCracha, &aniv.AniversarioEmpresa, &aniv.URLAniversarioEmpresaTv)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		fmt.Printf("Nome: %s\nAniversario empresa: %s\n URL: %s\n", aniv.NomeCracha, aniv.AniversarioEmpresa, aniv.URLAniversarioEmpresaTv)
		aniversariantes = append(aniversariantes, aniv)
	}

	// Return the list of employees and no error.
	return aniversariantes, nil

}
