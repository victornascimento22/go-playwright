package services

import (
	"bytes"
	"fmt"
	"time"

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
	//CaptureScreenshotService(body models.RequestBody) (screenshot []byte, err error)
	CaptureScreenshotServicePBI(body models.RequestBody) (screenshot []byte, err error)
	CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error)
}

// ScreenshotService implementa a interface IScreenshotService
type ScreenshotService struct {
}

// ConstructorScreenshotService creates a new instance of ScreenshotService
// Returns:
//   - IScreenshotService: New instance of the service
func ConstructorScreenshotService() *ScreenshotService {
	return &ScreenshotService{}
}

// Metodo da interface
func (s *ScreenshotService) CaptureScreenshotServicePBI(body models.RequestBody) (screenshot []byte, err error) {
	return CaptureScreenshotServicePBI(body)
}

// metodo da interface
func (s *ScreenshotService) CaptureScreenshotServiceGeneric(body models.RequestBody) (screenshot []byte, err error) {
	return CaptureScreenshotServiceGeneric(body)
}

// CaptureScreenshotServicePBI implementação específica para Power BI (stub)
func CaptureScreenshotServicePBI(body models.RequestBody) (screenshot []byte, err error) {
	// Inicializa o navegador
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Configura a página com viewport maior para melhor qualidade
	page := browser.MustPage(body.Url).MustSetViewport(1920, 1080, 4, false)

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
