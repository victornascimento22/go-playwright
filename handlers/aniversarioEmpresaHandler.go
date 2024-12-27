package handler

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

type IAniversariantesEmpresaController struct {
	controller controller.IAniversarioEmpresaController
}

func ConstructorGetAniversarioEmpresaController(ctrl controller.IAniversarioEmpresaController) *IAniversariantesEmpresaController {
	return &IAniversariantesEmpresaController{
		controller: ctrl,
	}
}

func (h *IAniversariantesEmpresaController) Handler(c *gin.Context) {
	aniversariantes, err := h.controller.GetAniversarioEmpresaController()

	if err != nil {
		// Verifique o tipo de erro para ajustar o código de status se necessário
		c.JSON(500, gin.H{"error": "Erro ao buscar aniversariantes: " + err.Error()})
		return
	}

	if len(aniversariantes) == 0 {
		// Retorne 204 (No Content) se não houver dados
		c.JSON(204, gin.H{"message": "Nenhum aniversariante encontrado"})
		return
	}

	c.JSON(200, aniversariantes)
}
