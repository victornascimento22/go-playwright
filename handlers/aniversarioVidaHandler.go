package handlers

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

type IAniversariantesVidaHandler interface {
	GetAniversariantesVidaHandler(c *gin.Context)
}
type AniversariantesVidaController struct {
	controller controller.IAniversariantesVidaController
}

func ConstructorAniversariantesVidaController(c controller.IAniversariantesVidaController) *AniversariantesVidaController {
	return &AniversariantesVidaController{
		controller: c,
	}
}

func (h *AniversariantesVidaController) GetAniversariantesVidaHandler(c *gin.Context) {
	aniversariantes, err := h.controller.GetAniversariantesVidaController()
	if err != nil {
		// Respond with a 500 status code and an error message if an error occurs.
		c.JSON(500, gin.H{"error": "Erro ao buscar aniversariantes de vida: " + err.Error()})
		return
	}

	if len(aniversariantes) == 0 {
		// Respond with a 204 status code if no anniversaries are found.
		c.JSON(200, gin.H{"message": "Nenhum aniversariante de vida encontrado hoje"})
		return
	}

	// Respond with a 200 status code and the list of anniversaries if successful.
	c.JSON(200, aniversariantes)
}
