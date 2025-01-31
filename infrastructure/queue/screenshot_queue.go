package infra

import (
	"log"

	"gitlab.com/applications2285147/api-go/internal/models"
)

type IScreenshotQueue struct {
	requests chan models.ScreenshotRequest
	service  IScreenshotProcessor // interface para processar
}

// Adicione esta interface
type IScreenshotProcessor interface {
	CaptureScreenshotServicePBI(url string) ([]byte, error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) ([]byte, error)
	SendToRaspberry(screenshot []byte, raspberryIP string) error
	EnqueueScreenshot(url string, raspberryIP string, isPBI bool)
}

func NewScreenshotQueue(service IScreenshotProcessor) *IScreenshotQueue {
	sq := &IScreenshotQueue{
		requests: make(chan models.ScreenshotRequest, 100),
		service:  service,
	}
	// Inicia o worker
	go sq.processRequests()
	return sq
}

func (sq *IScreenshotQueue) AddRequest(request models.ScreenshotRequest) {
	log.Printf("📥 Adicionando request na fila para TV: %s", request.RaspberryIP)
	sq.requests <- request
}

func (sq *IScreenshotQueue) GetRequests() chan models.ScreenshotRequest {
	return sq.requests
}

func (sq *IScreenshotQueue) processRequests() {
	for request := range sq.requests {
		log.Printf("⚙️ Processando da fila: TV %s, URL %s", request.RaspberryIP, request.URL)

		var screenshot []byte
		var err error

		// Processa baseado no tipo
		if request.IsPBI {
			screenshot, err = sq.service.CaptureScreenshotServicePBI(request.URL)
		} else {
			screenshot, err = sq.service.CaptureScreenshotServiceGeneric(models.RequestBody{
				Url: request.URL,
			})
		}

		if err != nil {
			log.Printf("❌ Erro processando screenshot: %v", err)
			continue
		}

		// Envia para Raspberry
		if err := sq.service.SendToRaspberry(screenshot, request.RaspberryIP); err != nil {
			log.Printf("❌ Erro enviando para Raspberry: %v", err)
			continue
		}

		log.Printf("✅ Processamento concluído para TV %s", request.RaspberryIP)
	}
}

func (sq *IScreenshotQueue) EnqueueScreenshot(url string, raspberryIP string, isPBI bool) {
	log.Printf("🔄 Enfileirando screenshot para TV: %s", raspberryIP)
	request := models.ScreenshotRequest{
		URL:         url,
		RaspberryIP: raspberryIP,
		IsPBI:       isPBI,
	}
	sq.AddRequest(request)
	log.Printf("✨ Request enfileirada com sucesso para TV: %s", raspberryIP)
}
