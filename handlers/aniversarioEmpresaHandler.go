// Package handler provides HTTP handlers for managing API endpoints.
package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

type IAniversariantesEmpresaHandler interface {
	GetAniversariantesEmpresaHandler(c *gin.Context)
}

// IAniversariantesEmpresaController encapsulates the logic for handling anniversary-related requests.
type IAniversariantesEmpresaController struct {
	// controller provides the business logic for retrieving anniversaries.
	controller controller.IAniversarioEmpresaController
}

// ConstructorGetAniversarioEmpresaController creates a new instance of IAniversariantesEmpresaController.
// ctrl: an implementation of IAniversarioEmpresaController.
// Returns a pointer to IAniversariantesEmpresaController.
func ConstructorGetAniversarioEmpresaController(ctrl controller.IAniversarioEmpresaController) *IAniversariantesEmpresaController {
	return &IAniversariantesEmpresaController{
		controller: ctrl,
	}
}

// Handler processes HTTP requests for retrieving employee anniversaries.
// c: the Gin context for the HTTP request.
func (h *IAniversariantesEmpresaController) GetAniversariantesEmpresaHandler(c *gin.Context) {

	log.Println("Iniciando busca de aniversariantes")

	aniversariantes, err := h.controller.GetAniversarioEmpresaController()
	if err != nil {
		log.Printf("Erro ao buscar aniversariantes: %v", err)
		c.JSON(500, gin.H{
			"error":   "Erro ao buscar aniversariantes de empresa",
			"details": err.Error(),
		})
		return
	}

	if len(aniversariantes) == 0 {
		c.JSON(200, gin.H{
			"message": "Nenhum aniversariante de empresa encontrado hoje",
		})
		return
	}

	c.JSON(200, aniversariantes)
}
