// Package handler provides HTTP handlers for managing API endpoints.
package handler

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

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
func (h *IAniversariantesEmpresaController) Handler(c *gin.Context) {
	// Call the controller to fetch anniversaries.
	aniversariantes, err := h.controller.GetAniversarioEmpresaController()

	if err != nil {
		// Respond with a 500 status code and an error message if an error occurs.
		c.JSON(500, gin.H{"error": "Erro ao buscar aniversariantes: " + err.Error()})
		return
	}

	if len(aniversariantes) == 0 {
		// Respond with a 204 status code if no anniversaries are found.
		c.JSON(204, gin.H{"message": "Nenhum aniversariante encontrado"})
		return
	}

	// Respond with a 200 status code and the list of anniversaries if successful.
	c.JSON(200, aniversariantes)
}
