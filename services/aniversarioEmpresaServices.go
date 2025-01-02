// Package services contains the business logic for application operations.
package services

import (
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/internal/models"
)

// IAniversarioEmpresaServices defines the interface for anniversary-related services.
type IAniversarioEmpresaServices interface {
	// GetAniversariantesEmpresaService retrieves the list of employee anniversaries.
	GetAniversariantesEmpresaService() ([]models.Aniversariantes, error)
}

// IAniversarioEmpresaRepositorys provides the implementation of IAniversarioEmpresaServices
// and acts as a bridge to the repository layer.
type IAniversarioEmpresaRepositorys struct {
	// repo is the repository used to access anniversary data.
	repo repository.IAniversariantesEmpresaRepository
}

// ConstructorIAniversarioEmpresaRepositorys creates a new instance of IAniversarioEmpresaRepositorys.
// i: an implementation of IAniversariantesEmpresaRepository.
// Returns a pointer to IAniversarioEmpresaRepositorys.
func ConstructorIAniversarioEmpresaRepositorys(i repository.IAniversariantesEmpresaRepository) *IAniversarioEmpresaRepositorys {
	return &IAniversarioEmpresaRepositorys{
		repo: i,
	}
}

// GetAniversariantesEmpresaService calls the repository method to fetch employee anniversaries
// and applies any necessary business logic.
// Returns a slice of Aniversariantes models or an error if the operation fails.
func (a *IAniversarioEmpresaRepositorys) GetAniversariantesEmpresaService() ([]models.Aniversariantes, error) {
	// Call the repository to fetch the list of anniversaries.
	aniversariantes, err := a.repo.BuscarAniversariantesEmpresa()
	if err != nil {
		return nil, err
	}

	// Additional business logic can be added here if required.
	return aniversariantes, nil
}
