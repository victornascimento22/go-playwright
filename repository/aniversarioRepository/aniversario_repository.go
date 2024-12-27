/
package repository

import (
	"database/sql"
	"fmt"

	"gitlab.com/applications2285147/api-go/internal/models"
)

type repository struct {
	db *sql.DB
}

func NewAniversarioRepository(db *sql.DB) *AniversarioRepository {
	return &AniversarioRepository{
		db: db,
	}
}

func (r *AniversarioRepository) BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error) {
	// Query para buscar aniversariantes do dia
	query := `SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv
		FROM DADOS_FUNCIONARIOS
		WHERE date_part('day', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('day', CURRENT_DATE)
		AND date_part('month', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('month', CURRENT_DATE);`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aniversariantes []models.Aniversariantes

	for rows.Next() {
		var aniv models.Aniversariantes
		err := rows.Scan(&aniv.Nome_cracha, &aniv.Aniversario_empresa, &aniv.Url_aniversario_empresa_tv)
		fmt.Printf("Nome: %s\nAniversario empresa: %s\n URL: %s\n", aniv.Nome_cracha, aniv.Aniversario_empresa, aniv.Url_aniversario_empresa_tv)
		if err != nil {
			return nil, err
		}
		aniversariantes = append(aniversariantes, aniv)
	}

	return []aniversariantes, nil

}
