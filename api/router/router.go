// Package api provides the routing
package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
func (r *RouterHandler) SetupRouter() *gin.Engine {
	router := gin.Default()
	// Configuração CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configuração de rotas
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Rotas agrupadas
	aniversario := router.Group("/aniversario")
	{
		aniversario.GET("/getAniversariosEmpresa", r.aniversariosHandler.empresaHandler.GetAniversariantesEmpresaHandler)
		aniversario.GET("/getAniversariosVida", r.aniversariosHandler.vidaHandler.GetAniversariantesVidaHandler)
	}

	screenshots := router.Group("/screenshots")
	{
		screenshots.POST("/update", r.screenshotHandler.screenshotHandler.UpdateScreenshotHandler)
	}

	return router
}
