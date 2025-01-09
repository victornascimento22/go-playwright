package controller

import (
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

type IScreenshotController interface {
	PostScreenshotController(url *string) ([]byte, error)
	CaptureScreenshotServicePBI(url string) ([]byte, error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) ([]byte, error)
	SendToRaspberry(screenshot []byte, raspberryIP string) error
}

type IScreenshotServices struct {
	services services.IScreenshotService
}

func ConstructorIScreenshotServices(services services.IScreenshotService) *IScreenshotServices {
	return &IScreenshotServices{
		services: services,
	}
}

func (x *IScreenshotServices) PostScreenshotController(url *string) (screenshot []byte, err error) {
	return x.services.CaptureScreenshotService(models.RequestBody{Url: *url})
}

func (c *IScreenshotServices) CaptureScreenshotServicePBI(url string) ([]byte, error) {
	return c.services.CaptureScreenshotServicePBI(url)
}

func (c *IScreenshotServices) CaptureScreenshotServiceGeneric(body models.RequestBody) ([]byte, error) {
	return c.services.CaptureScreenshotServiceGeneric(body)
}

func (c *IScreenshotServices) SendToRaspberry(screenshot []byte, raspberryIP string) error {
	return c.services.SendToRaspberry(screenshot, raspberryIP)
}
