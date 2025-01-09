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
	"gitlab.com/applications2285147/api-go/internal/models"
)

// IScreenshotService define a interface para serviços de captura de screenshots
type IScreenshotService interface {
	// CaptureScreenshotService captura um screenshot de uma URL fornecida
	// Params:
	//   - body: Contém a URL e outros parâmetros necessários para a captura
	// Returns:
	//   - screenshot: Imagem em bytes
	//   - err: Erro, se houver
	CaptureScreenshotService(body models.RequestBody) (screenshot []byte, err error)
	CaptureScreenshotServicePBI(url string) (screenshot []byte, err error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error)
	SendToRaspberry(screenshot []byte, raspberryIP string) error
	SendMultipleToRaspberry(config models.DisplayConfig) error
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

// ScreenshotService implementa a interface IScreenshotService
type ScreenshotService struct{}

// ConstructorScreenshotService cria uma nova instância de ScreenshotService
// Returns:
//   - IScreenshotService: Nova instância do serviço
func ConstructorScreenshotService() IScreenshotService {
	return &ScreenshotService{}
}

// CaptureScreenshotService implementa o método da interface usando a função genérica
// Params:
//   - body: Contém a URL e parâmetros para captura
//
// Returns:
//   - screenshot: Imagem capturada em bytes
//   - err: Erro, se houver
func (s *ScreenshotService) CaptureScreenshotService(body models.RequestBody) (screenshot []byte, err error) {
	return CaptureScreenshotServiceGeneric(body)
}

func (s *ScreenshotService) SendToRaspberry(screenshot []byte, raspberryIP string) error {
	url := fmt.Sprintf("http://%s:8081/webhook", raspberryIP)

	// Codifica a imagem em base64
	base64Image := base64.StdEncoding.EncodeToString(screenshot)

	payload := struct {
		Image          string `json:"image"` // Mudamos para string (base64)
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

func (s *ScreenshotService) SendMultipleToRaspberry(config models.DisplayConfig) error {
	// Captura screenshots de todas as URLs
	for i, urlConfig := range config.URLs {
		var screenshot []byte
		var err error

		// Escolhe o método baseado no source de cada URL
		if urlConfig.Source == "pbi" {
			screenshot, err = CaptureScreenshotServicePBI(urlConfig.URL)
		} else {
			screenshot, err = CaptureScreenshotServiceGeneric(models.RequestBody{
				Url: urlConfig.URL,
			})
		}

		if err != nil {
			return fmt.Errorf("erro ao capturar screenshot %d: %v", i+1, err)
		}

		// Codifica a imagem em base64
		base64Image := base64.StdEncoding.EncodeToString(screenshot)

		// Envia cada screenshot com seu índice
		url := fmt.Sprintf("http://%s:8081/webhook", config.RaspberryIP)
		payload := struct {
			Image          string `json:"image"` // Mudamos para string (base64)
			Index          int    `json:"index"`
			TransitionTime int    `json:"transition_time"`
		}{
			Image:          base64Image,
			Index:          i,
			TransitionTime: config.TransitionTime,
		}

		// Envia para o webhook
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("erro ao criar payload %d: %v", i+1, err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
		if err != nil {
			return fmt.Errorf("erro ao enviar screenshot %d: %v", i+1, err)
		}
		defer resp.Body.Close()

		// Lê resposta
		body, _ := io.ReadAll(resp.Body)
		log.Printf("✅ Screenshot %d enviado. Resposta: %s", i+1, string(body))
	}
	return nil
}

func (s *ScreenshotService) CaptureScreenshotServicePBI(url string) (screenshot []byte, err error) {
	return CaptureScreenshotServicePBI(url)
}

func (s *ScreenshotService) CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error) {
	return CaptureScreenshotServiceGeneric(body)
}
