package handlers

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller"
)

type IScreenshotHandler interface {
	UpdateScreenshotHandler(c *gin.Context)
}

type IScreenshotController struct {
	controller controller.IScreenshotController
}

func ConstructorIScreenshotController(ctrl controller.IScreenshotController) *IScreenshotController {
	return &IScreenshotController{
		controller: ctrl,
	}
}

func (h *IScreenshotController) UpdateScreenshotHandler(c *gin.Context) {
	// Call the controller method to handle the request
	h.controller.UpdateScreenshotController(c)
}
