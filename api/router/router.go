// Package api provides the routing
package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	handler "gitlab.com/applications2285147/api-go/handlers"
)

// IAniversariantesEmpresaController encapsulates the controller logic for handling anniversary-related requests.

type IAniversariantesHandler struct {
	empresaHandler handler.IAniversariantesEmpresaHandler
	vidaHandler    handler.IAniversariantesVidaHandler
}
type IScreenshotHandler struct {
	screenshotHandler handler.IScreenshotHandler
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

func ConstructorScreenshotHandler(handScreenshot handler.IScreenshotHandler) *IScreenshotHandler {
	return &IScreenshotHandler{
		screenshotHandler: handScreenshot,
	}
}

type RouterHandler struct {
	aniversariosHandler *IAniversariantesHandler
	screenshotHandler   *IScreenshotHandler
}

func NewRouterHandler(a *IAniversariantesHandler, s *IScreenshotHandler) *RouterHandler {
	return &RouterHandler{
		aniversariosHandler: a,
		screenshotHandler:   s,
	}
}

// Router sets up the HTTP routes for the anniversary-related endpoints.
// db: Database connection used by the application.
func (r *RouterHandler) Router(db *sql.DB) {
	// Initialize the Gin router.
	router := gin.Default()

	// Serve arquivos estáticos
	router.LoadHTMLGlob("templates/*")

	// Rota para o frontend
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Define a route group for anniversary-related endpoints.
	aniversario := router.Group("/aniversario")
	{
		// Define a GET endpoint to retrieve employee anniversaries.
		aniversario.GET("/getAniversariosEmpresa", func(c *gin.Context) {
			r.aniversariosHandler.empresaHandler.GetAniversariantesEmpresaHandler(c)
		})

		aniversario.GET("/getAniversariosVida", func(c *gin.Context) {
			r.aniversariosHandler.vidaHandler.GetAniversariantesVidaHandler(c)
		})
	}

	screenshots := router.Group("/screenshots")
	{
		screenshots.POST("/getScreenshot", func(c *gin.Context) {
			r.screenshotHandler.screenshotHandler.PostScreenshotHandler(c)
		})
		screenshots.POST("/update", func(c *gin.Context) {
			r.screenshotHandler.screenshotHandler.UpdateDisplayHandler(c)
		})
	}

	// Start the HTTP server on port 8080.
	router.Run(":8080")
}
