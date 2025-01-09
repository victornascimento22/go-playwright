package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/screenshotControllers"
	"gitlab.com/applications2285147/api-go/internal/models"
)

type IScreenshotHandler interface {
	PostScreenshotHandler(c *gin.Context)
	UpdateDisplayHandler(c *gin.Context)
}

type IScreenshotController struct {
	controller controller.IScreenshotController
}

func ConstructorScreenshotController(ctrl controller.IScreenshotController) *IScreenshotController {
	return &IScreenshotController{
		controller: ctrl,
	}
}

func (x *IScreenshotController) PostScreenshotHandler(c *gin.Context) {
	var body models.RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	screenshot, err := x.controller.PostScreenshotController(&body.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"screenshot": screenshot})
}

func (h *IScreenshotController) UpdateDisplayHandler(c *gin.Context) {
	var req struct {
		URLs           []models.URLConfig `json:"urls"`
		TransitionTime int                `json:"transition_time"`
		RaspberryIP    string             `json:"raspberry_ip"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Processa cada URL
	for _, urlConfig := range req.URLs {
		var screenshot []byte
		var err error

		if urlConfig.Source == "pbi" {
			screenshot, err = h.controller.CaptureScreenshotServicePBI(urlConfig.URL)
		} else {
			screenshot, err = h.controller.CaptureScreenshotServiceGeneric(models.RequestBody{
				Url: urlConfig.URL,
			})
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao capturar screenshot: " + err.Error()})
			return
		}

		if err := h.controller.SendToRaspberry(screenshot, req.RaspberryIP); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar para Raspberry: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Display atualizado com sucesso"})
}
