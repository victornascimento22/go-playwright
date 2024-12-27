package controller

import (
	"gitlab.com/applications2285147/api-go/database"
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/internal/models"
)

func GetAniversarioEmpresaController() ([]models.Aniversariantes, error) {
	repo := repository.ConstructorAniversarioEmpresa(database.DB)
	return repo.BuscarAniversariantesEmpresa()
}
