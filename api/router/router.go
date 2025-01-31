// Package api provides the routing
package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/applications2285147/api-go/handlers"
	"gitlab.com/applications2285147/api-go/services"
)

// IAniversariantesEmpresaController encapsulates the controller logic for handling anniversary-related requests.
type IRouter interface {
	SetupRouter() (*gin.Engine, error)
}

type IAniversariantesHandler struct {
	AniversarioEmpresaHandler handlers.IAniversariantesEmpresaHandler
}

type IScreenshotHandler struct {
	ScreenshotHandler handlers.IScreenshotHandler
}

type IWS struct {
	WebsocketHandler *handlers.WebsocketHandlerImpl
}

// RouterHandler manages the routes and their handlers.
type RouterHandler struct {
	AniversariosHandler handlers.IAniversariantesEmpresaHandler
	ScreenshotHandler   handlers.IScreenshotHandler
	WebsocketHandler    *IWS
	ScreenshotService   services.IScreenshotService
}

// NewRouterHandler creates a new RouterHandler instance.
func NewRouterHandler(a handlers.IAniversariantesEmpresaHandler, s handlers.IScreenshotHandler, w *IWS, ss services.IScreenshotService) *RouterHandler {
	return &RouterHandler{
		AniversariosHandler: a,
		ScreenshotHandler:   s,
		WebsocketHandler:    w,
		ScreenshotService:   ss,
	}
}

// Router sets up the HTTP routes for the anniversary-related endpoints.
// db: Database connection used by the application.
func (r *RouterHandler) SetupRouter() (*gin.Engine, error) {
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
	// Rotas agrupadas
	aniversario := router.Group("/aniversario")
	{
		aniversario.GET("/getAniversariosEmpresa", r.AniversariosHandler.GetAniversariantesEmpresaHandler)
	}

	screenshots := router.Group("/screenshots")
	{
		screenshots.POST("/update", r.ScreenshotHandler.UpdateScreenshotHandler)
	}

	websockets := router.Group("/ws")
	{
		websockets.GET("/connect", r.WebsocketHandler.WebsocketHandler.WebsocketHandler)
	}
	return router, nil
}

// AniversariantesHandler implements the IAniversariantesEmpresaHandler interface.
type AniversariantesHandler struct {
	EmpresaHandler handlers.IAniversariantesEmpresaHandler
	VidaHandler    handlers.IAniversariantesVidaHandler
}

// Ensure that AniversariantesHandler implements the IAniversariantesEmpresaHandler interface
func (h *AniversariantesHandler) GetAniversariantesEmpresaHandler(c *gin.Context) {
	h.EmpresaHandler.GetAniversariantesEmpresaHandler(c)
}
