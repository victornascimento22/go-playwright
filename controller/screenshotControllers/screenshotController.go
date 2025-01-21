package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

// IScreenshotController define a interface para o controlador de captura de tela
type IScreenshotController interface {
	// Método para postar captura de tela// Método para captura genérica// Método para enviar captura para Raspberry Pi
	UpdateScreenshotController(c *gin.Context)
}

// IScreenshotServices estrutura que contém os serviços de captura de tela
type IScreenshotServices struct {
	services services.IScreenshotService // Referência aos serviços de captura de tela
}

// ConstructorIScreenshotServices cria uma nova instância de IScreenshotServices
func ConstructorIScreenshotServices(services services.IScreenshotService) *IScreenshotServices {
	return &IScreenshotServices{
		services: services, // Inicializa a estrutura com os serviços fornecidos
	}
}

// PostScreenshotController implementa o método para postar captura de tela
func (x *IScreenshotServices) UpdateScreenshotController(c *gin.Context) {
	var req struct {
		URLs           []models.URLConfig `json:"urls"`
		TransitionTime int                `json:"transition_time"`
		RaspberryIP    string             `json:"raspberry_ip"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Para cada URL
	for _, urlConfig := range req.URLs {
		// Enfileira sem checar erro
		x.services.EnqueueScreenshot(
			urlConfig.URL,
			req.RaspberryIP,
			urlConfig.Source == "pbi",
		)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Screenshots enfileiradas com sucesso",
		"count":   len(req.URLs),
	})
}
