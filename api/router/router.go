// Package api provides the routing
package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

// IAniversariantesEmpresaController encapsulates the controller logic for handling anniversary-related requests.
type IAniversariantesEmpresaController struct {
	controller controller.IAniversarioEmpresaController // Interface for the business logic layer.
}

// ConstructorGetAniversarioEmpresaController initializes the controller structure.
// ctrl: An implementation of IAniversarioEmpresaController.
// Returns an instance of IAniversariantesEmpresaController.
func ConstructorGetAniversarioEmpresaController(ctrl controller.IAniversarioEmpresaController) *IAniversariantesEmpresaController {
	return &IAniversariantesEmpresaController{
		controller: ctrl,
	}
}

// Router sets up the HTTP routes for the anniversary-related endpoints.
// db: Database connection used by the application.
func (x *IAniversariantesEmpresaController) Router(db *sql.DB) {
	// Initialize the Gin router.
	r := gin.Default()

	// Define a route group for anniversary-related endpoints.
	aniversario := r.Group("/aniversario")
	{
		// Define a GET endpoint to retrieve employee anniversaries.
		aniversario.GET("/getAniversariosEmpresa", func(c *gin.Context) {
			// Call the controller to fetch anniversary data.
			aniversariantes, err := x.controller.GetAniversarioEmpresaController()
			if err != nil {
				// Respond with a 500 Internal Server Error if fetching data fails.
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Respond with a 200 OK status and the data if fetching is successful.
			c.JSON(http.StatusOK, aniversariantes)
		})
	}

	// Start the HTTP server on port 8080.
	r.Run(":8080")
}
