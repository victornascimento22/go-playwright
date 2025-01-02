package controller

import (
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

type IAniversarioEmpresaController interface {
	GetAniversarioEmpresaController() ([]models.Aniversariantes, error)
}

type IAniversarioEmpresaServices struct {
	services services.IAniversarioEmpresaServices
}

func ConstructorIAniversarianteEmpresaServices(services services.IAniversarioEmpresaServices) *IAniversarioEmpresaServices {
	return &IAniversarioEmpresaServices{
		services: services,
	}
}

func (x *IAniversarioEmpresaServices) GetAniversarioEmpresaController() ([]models.Aniversariantes, error) {
	aniversariantes, err := x.services.GetAniversariantesEmpresaService()
	if err != nil {
		panic("OI")
	}
	return aniversariantes, err
}
