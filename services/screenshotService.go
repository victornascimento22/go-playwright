package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"encoding/base64"

	"image"

	"github.com/disintegration/imaging"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"gitlab.com/applications2285147/api-go/infrastructure/queue"
	"gitlab.com/applications2285147/api-go/internal/models"
	queueModels "gitlab.com/applications2285147/api-go/internal/queue"
)

// IScreenshotService define a interface para serviços de captura de screenshots
type IScreenshotService interface {
	// CaptureScreenshotService captura um screenshot de uma URL fornecida
	// Params:
	//   - body: Contém a URL e outros parâmetros necessários para a captura
	// Returns:
	//   - screenshot: Imagem em bytes
	//   - err: Erro, se houver
	//CaptureScreenshotService(body models.RequestBody) (screenshot []byte, err error)
	CaptureScreenshotServicePBI(url string) (screenshot []byte, err error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error)
	SendToRaspberry(screenshot []byte, raspberryIP string) error
	EnqueueScreenshot(url string, raspberryIP string, isPBI bool)
}

// ScreenshotService implementa a interface IScreenshotService
type ScreenshotService struct {
	services  IScreenshotService
	queue     *queue.ScreenshotQueue
	wsManager *WebSocketManager
}

// ConstructorScreenshotService creates a new instance of ScreenshotService
// Returns:
//   - IScreenshotService: New instance of the service
func ConstructorScreenshotService(x IScreenshotService, queue *queue.ScreenshotQueue, wsManager *WebSocketManager) IScreenshotService {
	return &ScreenshotService{
		services:  x,
		queue:     queue,
		wsManager: wsManager,
	}
}

// Metodo da interface
func (s *ScreenshotService) CaptureScreenshotServicePBI(url string) (screenshot []byte, err error) {
	return CaptureScreenshotServicePBI(url)
}

// metodo da interface
func (s *ScreenshotService) CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error) {
	return CaptureScreenshotServiceGeneric(body)
}

// CaptureScreenshotServicePBI implementação específica para Power BI (stub)
func CaptureScreenshotServicePBI(url string) (screenshot []byte, err error) {
	// Inicializa o navegador
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Configura a página com viewport maior para melhor qualidade
	page := browser.MustPage(url).MustSetViewport(1920, 1080, 4, false)

	// Aguarda o carregamento inicial
	page.MustWaitLoad()
	time.Sleep(15 * time.Second)

	// Remove elementos indesejados (barra de status e logo)
	page.MustEval(`() => {
		const statusBar = document.querySelector("#reportLandingContainer > div > exploration-container > div > div > docking-container > div > pbi-status-bar > section");
		if (statusBar) statusBar.style.display = "none";

		const logo = document.querySelector("#embedWrapperID > div.logoBarWrapper > logo-bar > div > div");
		if (logo) logo.style.display = "none";
	}`)

	// Aguarda as alterações serem aplicadas
	time.Sleep(2 * time.Second)

	// Captura a screenshot
	fullScreenshot, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  1920,
			Height: 1080,
			Scale:  1,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao capturar screenshot: %v", err)
	}

	// Processa a imagem para cortar as bordas
	img, err := imaging.Decode(bytes.NewReader(fullScreenshot))
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar imagem: %v", err)
	}

	// Obtém as dimensões
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	// Define as margens para corte
	left := 50
	top := 0
	right := width - 50
	bottom := height - 20

	// Corta a imagem
	croppedImg := imaging.Crop(img, image.Rect(left, top, right, bottom))

	// Converte a imagem processada de volta para bytes
	var buf bytes.Buffer
	if err := imaging.Encode(&buf, croppedImg, imaging.PNG); err != nil {
		return nil, fmt.Errorf("erro ao codificar imagem: %v", err)
	}

	return buf.Bytes(), nil
}

// CaptureScreenshotServiceGeneric implementa a captura de screenshot genérica usando rod
// Params:
//   - body: Contém a URL alvo para captura
//
// Returns:
//   - screenshot: Imagem capturada em bytes
//   - err: Erro, se houver
func CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(body.Url).MustSetViewport(1920, 1080, 1, false)
	page.MustWaitLoad()

	time.Sleep(15 * time.Second) // Simula tempo de espera para carregamento completo
	screenshot, err = page.Screenshot(true, nil)
	if err != nil {
		return nil, err
	}
	return screenshot, nil
}

// CaptureScreenshotService implementa o método da interface usando a função genérica
// Params:
//   - body: Contém a URL e parâmetros para captura
//
// Returns:
//   - screenshot: Imagem capturada em bytes
//   - err: Erro, se houver

func (s *ScreenshotService) SendToRaspberry(screenshot []byte, raspberryIP string) error {
	// Tenta primeiro via WebSocket
	if conn, exists := s.wsManager.GetConnection(raspberryIP); exists {
		log.Printf("📡 Enviando via WebSocket para %s", raspberryIP)
		return conn.WriteJSON(map[string]interface{}{
			"image":           base64.StdEncoding.EncodeToString(screenshot),
			"index":           0,
			"transition_time": 15,
		})
	}

	// Fallback para HTTP se WebSocket não disponível
	log.Printf("📤 Enviando via HTTP para %s", raspberryIP)
	url := fmt.Sprintf("http://%s:8081/webhook", raspberryIP)

	// Seu código HTTP existente
	base64Image := base64.StdEncoding.EncodeToString(screenshot)
	payload := struct {
		Image          string `json:"image"`
		Index          int    `json:"index"`
		TransitionTime int    `json:"transition_time"`
	}{
		Image:          base64Image,
		Index:          0,
		TransitionTime: 15,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("erro ao criar payload: %v", err)
	}

	log.Printf("📤 Enviando payload para %s\n", url)
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Printf("❌ Erro ao enviar para raspberry: %v\n", err)
		return fmt.Errorf("erro ao enviar para raspberry: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("✅ Resposta do raspberry: %s (status: %d)\n", string(body), resp.StatusCode)
	return nil
}

// Adicione este método na sua struct ScreenshotService
func (s *ScreenshotService) EnqueueScreenshot(url string, raspberryIP string, isPBI bool) {
	log.Printf("🔄 Enfileirando screenshot para TV: %s", raspberryIP)
	request := queueModels.ScreenshotRequest{
		URL:         url,
		RaspberryIP: raspberryIP,
		IsPBI:       isPBI,
	}
	s.queue.AddRequest(request)
	log.Printf("✨ Request enfileirada com sucesso para TV: %s", raspberryIP)
}
