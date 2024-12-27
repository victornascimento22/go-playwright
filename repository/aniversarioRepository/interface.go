package repository

import (
	"gitlab.com/applications2285147/api-go/internal/models"
)

type AniversarioRepository interface {
	BuscarAniversariantesEmpresa() ([]models.Aniversariantes, error)
}
