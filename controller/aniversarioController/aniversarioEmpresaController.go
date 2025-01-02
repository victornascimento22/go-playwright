// Package controller handles the business logic and control flow of birthday-related operations
package controller

import (
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

// IAniversarioEmpresaController defines the interface for company birthday operations
type IAniversarioEmpresaController interface {
	// GetAniversarioEmpresaController retrieves a list of company birthdays
	GetAniversarioEmpresaController() ([]models.Aniversariantes, error)
}

// IAniversarioEmpresaServices implements the company birthday controller operations
type IAniversarioEmpresaServices struct {
	services services.IAniversarioEmpresaServices
}

// ConstructorIAniversarianteEmpresaServices creates a new instance of IAniversarioEmpresaServices
func ConstructorIAniversarianteEmpresaServices(services services.IAniversarioEmpresaServices) *IAniversarioEmpresaServices {
	return &IAniversarioEmpresaServices{
		services: services,
	}
}

// GetAniversarioEmpresaController retrieves company birthdays from the service layer
func (x *IAniversarioEmpresaServices) GetAniversarioEmpresaController() ([]models.Aniversariantes, error) {
	aniversariantes, err := x.services.GetAniversariantesEmpresaService()
	if err != nil {
		panic("OI")
	}
	return aniversariantes, err
}
