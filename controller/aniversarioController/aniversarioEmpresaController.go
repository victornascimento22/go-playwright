package controller

import (
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/internal/models"
)

type IAniversarioEmpresaController interface {
	GetAniversarioEmpresaController() ([]models.Aniversariantes, error)
}

type IAniversariantesEmpresaRepository struct {
	repository repository.IAniversariantesEmpresaRepository
}

func ConstructorIAniversarianteEmpresaRepository(rep repository.IAniversariantesEmpresaRepository) *IAniversariantesEmpresaRepository {
	return &IAniversariantesEmpresaRepository{
		repository: rep,
	}

}

func (x *IAniversariantesEmpresaRepository) GetAniversarioEmpresaController() ([]models.Aniversariantes, error) {

	repo, err := x.repository.BuscarAniversariantesEmpresa()

	if err != nil {
		panic("OI")
	}

	return repo, err

}
