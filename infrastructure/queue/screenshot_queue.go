package queue

import (
	"log"

	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/internal/queue"
)

type ScreenshotQueue struct {
	requests chan queue.ScreenshotRequest
	service  IScreenshotProcessor // interface para processar
}

// Adicione esta interface
type IScreenshotProcessor interface {
	CaptureScreenshotServicePBI(url string) ([]byte, error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) ([]byte, error)
	SendToRaspberry(screenshot []byte, raspberryIP string) error
}

func NewScreenshotQueue(service IScreenshotProcessor) *ScreenshotQueue {
	sq := &ScreenshotQueue{
		requests: make(chan queue.ScreenshotRequest, 100),
		service:  service,
	}
	// Inicia o worker
	go sq.processRequests()
	return sq
}

func (sq *ScreenshotQueue) AddRequest(request queue.ScreenshotRequest) {
	log.Printf("📥 Adicionando request na fila para TV: %s", request.RaspberryIP)
	sq.requests <- request
}

func (sq *ScreenshotQueue) GetRequests() chan queue.ScreenshotRequest {
	return sq.requests
}

func (sq *ScreenshotQueue) processRequests() {
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
