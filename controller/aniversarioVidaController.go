// Package controller fornece os controladores para gerenciamento de aniversariantes
package controller

import (
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

// IAniversariantesVidaController define a interface para operações relacionadas aos aniversariantes do seguro de vida
type IAniversariantesVidaController interface {
	GetAniversariantesVidaController() ([]models.Aniversariantes, error)
}

// IAniversariantesVidaServices implementa a estrutura do controlador de aniversariantes do seguro de vida
type IAniversariantesVidaServices struct {
	services services.IAniversariantesVidaServices
}

// ConstructorAniversariantesVidaServices cria uma nova instância do controlador de aniversariantes do seguro de vida
// Recebe services como dependência para acessar as operações de serviço
// Retorna um ponteiro para a estrutura IAniversariantesVidaServices inicializada
func ConstructorAniversariantesVidaServices(services services.IAniversariantesVidaServices) *IAniversariantesVidaServices {
	return &IAniversariantesVidaServices{
		services: services,
	}
}

// GetAniversariantesVidaController recupera a lista de aniversariantes do seguro de vida
// Retorna um slice de models.Aniversariantes e error
// Em caso de erro durante a busca, retorna nil e o erro encontrado
func (a *IAniversariantesVidaServices) GetAniversariantesVidaController() ([]models.Aniversariantes, error) {
	aniversariantes, err := a.services.GetAniversariantesVidaService()
	if err != nil {
		return nil, err
	}

	return aniversariantes, err
}
