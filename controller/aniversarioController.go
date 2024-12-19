// controller/aniversario_controller.go
package controller

import (
	"database/sql"

	"gitlab.com/applications2285147/api-go/internal/models"
)

type AniversarioController struct {
	db *sql.DB
}

func NewAniversarioController(db *sql.DB) *AniversarioController {
	return &AniversarioController{
		db: db,
	}
}

func (ac *AniversarioController) BuscarAniversariantesDoDia() ([]models.Aniversariantes, error) {
	// Query para buscar aniversariantes do dia
	query := `
        SELECT id, nome, data_nascimento 
        FROM aniversariantes 
        WHERE EXTRACT(MONTH FROM data_nascimento) = EXTRACT(MONTH FROM CURRENT_DATE)
        AND EXTRACT(DAY FROM data_nascimento) = EXTRACT(DAY FROM CURRENT_DATE)
    `

	rows, err := ac.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aniversariantes []models.Aniversariantes

	for rows.Next() {
		var aniv models.Aniversariantes
		err := rows.Scan(&aniv.ID, &aniv.Username, &aniv.Password)
		if err != nil {
			return nil, err
		}
		aniversariantes = append(aniversariantes, aniv)
	}

	return aniversariantes, nil
}
