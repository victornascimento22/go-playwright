package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra "gitlab.com/applications2285147/api-go/infrastructure/queue"
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

// IScreenshotController define a interface para o controlador de captura de tela
type IScreenshotController interface {
	// Método para postar captura de tela// Método para captura genérica// Método para enviar captura para Raspberry Pi
	UpdateScreenshotController(c *gin.Context)
}

// IScreenshotServices estrutura que con3tém os serviços de captura de tela
type IScreenshotServices struct {
	services services.IScreenshotService
	infra    infra.IScreenshotProcessor // Referência aos serviços de captura de tela
}

// ConstructorIScreenshotServices cria uma nova instância de IScreenshotServices
func ConstructorIScreenshotServices(services services.IScreenshotService) *IScreenshotServices {
	return &IScreenshotServices{
		services: services,
		// Inicializa a estrutura com os serviços fornecidos
	}
}

// PostScreenshotController implementa o método para postar captura de tela
func (x *IScreenshotServices) UpdateScreenshotController(c *gin.Context) {

	var req models.ScreenshotRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Para cada URL
	x.infra.EnqueueScreenshot(
		req.URL, // Passa todas as URLs de uma vez
		req.RaspberryIP,
		req.IsPBI,
	)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Screenshots enfileiradas com sucesso",
		"count":   len(req.URL),
	})
}
