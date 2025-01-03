// Package api provides the routing
package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handler "gitlab.com/applications2285147/api-go/handlers"
)

// IAniversariantesEmpresaController encapsulates the controller logic for handling anniversary-related requests.

type IAniversariantesHandler struct {
	empresaHandler handler.IAniversariantesEmpresaHandler
	vidaHandler    handler.IAniversariantesVidaHandler
}

// ConstructorGetAniversarioEmpresaController initializes the controller structure.
// ctrl: An implementation of IAniversarioEmpresaController.
// Returns an instance of IAniversariantesEmpresaController.
func ConstructorAniversariantesHandler(handEmpresa handler.IAniversariantesEmpresaHandler,
	handVida handler.IAniversariantesVidaHandler) *IAniversariantesHandler {
	return &IAniversariantesHandler{
		empresaHandler: handEmpresa,
		vidaHandler:    handVida,
	}
}

// Router sets up the HTTP routes for the anniversary-related endpoints.
// db: Database connection used by the application.
func (x *IAniversariantesHandler) Router(db *sql.DB) {
	// Initialize the Gin router.
	r := gin.Default()

	// Define a route group for anniversary-related endpoints.
	aniversario := r.Group("/aniversario")
	{
		// Define a GET endpoint to retrieve employee anniversaries.
		aniversario.GET("/getAniversariosEmpresa", func(c *gin.Context) {
			// Chama o handler diretamente, sem tentar capturar retornos
			x.empresaHandler.GetAniversariantesEmpresaHandler(c)
		})

		aniversario.GET("/getAniversariosVida", func(c *gin.Context) {
			x.vidaHandler.GetAniversariantesVidaHandler(c)
		})
	}

	// Start the HTTP server on port 8080.
	r.Run(":8080")
}
