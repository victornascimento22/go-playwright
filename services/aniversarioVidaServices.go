// Package services fornece a camada de serviços para manipulação de regras de negócio da aplicação
package services

import (
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/internal/models"
)

// IAniversariantesVidaServices define a interface para os serviços relacionados aos aniversariantes
// do produto Vida
type IAniversariantesVidaServices interface {
	GetAniversariantesVidaService() ([]models.Aniversariantes, error)
}

// IAniversariantesVidaRepositorys implementa a estrutura que contém o repositório
// para acesso aos dados dos aniversariantes do produto Vida
type IAniversariantesVidaRepositorys struct {
	rep repository.IAniversariantesVidaRepository
}

// ConstructorAniversariantesVidaRepositorys cria uma nova instância de IAniversariantesVidaRepositorys
// Parâmetros:
//   - repo: implementação da interface do repositório de aniversariantes
//
// Retorna:
//   - *IAniversariantesVidaRepositorys: ponteiro para a nova instância criada
func ConstructorAniversariantesVidaRepositorys(repo repository.IAniversariantesVidaRepository) *IAniversariantesVidaRepositorys {
	return &IAniversariantesVidaRepositorys{
		rep: repo,
	}
}

// GetAniversariantesVidaService recupera a lista de aniversariantes do produto Vida
// Este método faz a intermediação entre a camada de controller e repository
//
// Retorna:
//   - []models.Aniversariantes: slice contendo os dados dos aniversariantes
//   - error: possível erro ocorrido durante a operação
func (a *IAniversariantesVidaRepositorys) GetAniversariantesVidaService() ([]models.Aniversariantes, error) {
	aniversariantes, err := a.rep.GetAniversariantesVidaRepository()
	if err != nil {
		return nil, err
	}
	return aniversariantes, nil
}
