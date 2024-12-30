// controller/aniversario_controller.go
package repository

import (
	"fmt"

	infra "gitlab.com/applications2285147/api-go/infrastructure"
	"gitlab.com/applications2285147/api-go/internal/models"
)

type IAniversariantesEmpresaRepository interface {
	BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error)
}

type IConnectDatabase struct {
	infrastructure infra.IConnectDatabase
}

func ConstructorConnectDatabase(i infra.IConnectDatabase) *IConnectDatabase {

	return &IConnectDatabase{
		infrastructure: i,
	}

}

func (i *IConnectDatabase) BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error) {
	// Query para buscar aniversariantes do dia
	query := `SELECT nome_cracha, aniversario_empresa, url_aniversario_empresa_tv
		FROM DADOS_FUNCIONARIOS
		WHERE date_part('day', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('day', CURRENT_DATE)
		AND date_part('month', to_date(aniversario_empresa, 'DD/MM/YYYY')) = date_part('month', CURRENT_DATE);`
	db, err := i.infrastructure.ConnectDatabase()

	if err != nil {

		panic("oi")
	}

	rows, err := db.Query(query)
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

	return aniversariantes, nil

}
