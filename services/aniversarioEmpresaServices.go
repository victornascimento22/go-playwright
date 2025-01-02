package services

import (
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/internal/models"
)

type IAniversarioEmpresaServices interface {
	GetAniversariantesEmpresaService() ([]models.Aniversariantes, error)
}

type IAniversarioEmpresaRepositorys struct {
	repo repository.IAniversariantesEmpresaRepository
}

// Fixing the constructor to return the correct type
func ConstructorIAniversarioEmpresaRepositorys(i repository.IAniversariantesEmpresaRepository) *IAniversarioEmpresaRepositorys {
	return &IAniversarioEmpresaRepositorys{
		repo: i,
	}
}

// Nova função de serviço que chama o método
func (a *IAniversarioEmpresaRepositorys) GetAniversariantesEmpresaService() ([]models.Aniversariantes, error) {
	aniversariantes, err := a.repo.BuscarAniversariantesEmpresa()
	if err != nil {
		return nil, err
	}
	// Aqui você pode adicionar lógica adicional, se necessário
	return aniversariantes, nil
}
