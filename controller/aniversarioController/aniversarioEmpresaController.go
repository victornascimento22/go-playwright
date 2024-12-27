// controller/aniversario_controller.go
package controller

import "gitlab.com/applications2285147/api-go/internal/models"

func GetAniversarioEmpresaController() ([]models.Aniversariantes, error) {

	aniversariante, err := repository.BuscarAniversariantesEmpresa()
}
