package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gitlab.com/applications2285147/api-go/internal/models"
	"gitlab.com/applications2285147/api-go/services"
)

// Interface IWS com o método WebSocketHandler
type IWS interface {
	WebsocketHandler(c *gin.Context)
}

// DisplayConfig representa a configuração recebida via WebSocket
type DisplayConfig struct {
	URLs        []string `json:"urls"`
	Interval    int      `json:"interval"` // tempo em segundos
	RaspberryIP string   `json:"raspberryIP"`
	IsPBI       bool     `json:"isPBI"`
}

// WebsocketHandlerImpl gerencia as conexões WebSocket
type WebsocketHandlerImpl struct {
	clients           map[string]*websocket.Conn
	screenshotService services.IScreenshotService
}

// Constante para a porta dos clientes Raspberry
const RASPBERRY_PORT = "8081"

// Estrutura da mensagem
type DisplayMessage struct {
	Action   string   `json:"action"`   // "register", "display" ou "status"
	IP       string   `json:"ip"`       // IP do Raspberry
	URLs     []string `json:"urls"`     // URLs das imagens
	Interval int      `json:"interval"` // Tempo entre as imagens
}

// NewWebsocketHandler cria uma nova instância do handler
func NewWebsocketHandler(screenshotService services.IScreenshotService) *WebsocketHandlerImpl {
	return &WebsocketHandlerImpl{
		clients:           make(map[string]*websocket.Conn),
		screenshotService: screenshotService,
	}
}

// Configuração do upgrader para WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Adicionar método para verificar status
func (w *WebsocketHandlerImpl) IsClientConnected(ip string) bool {
	_, exists := w.clients[ip]
	return exists
}

// DisconnectClient desconecta um cliente específico
func (w *WebsocketHandlerImpl) DisconnectClient(ip string) {
	if conn, exists := w.clients[ip]; exists {
		conn.Close()
		delete(w.clients, ip)
		log.Printf("🔌 Cliente desconectado manualmente: %s:%s", ip, RASPBERRY_PORT)
	}
}

// WebsocketHandler gerencia a conexão WebSocket
func (w *WebsocketHandlerImpl) WebsocketHandler(c *gin.Context) {
	log.Println("Tentando fazer upgrade para WebSocket...")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ Erro ao conectar websocket: %v", err)
		return
	}
	log.Println("✅ Conexão WebSocket estabelecida com sucesso!")
	defer conn.Close()

	// Registrar o cliente automaticamente
	var initialMsg struct {
		Action string `json:"action"`
		IP     string `json:"ip"`
	}

	err = conn.ReadJSON(&initialMsg)
	if err != nil {
		log.Printf("❌ Erro ao ler mensagem inicial: %v", err)
		return
	}

	// Registrar o cliente usando o IP enviado pelo frontend
	w.clients[initialMsg.IP] = conn
	log.Printf("✅ Cliente registrado automaticamente: %s", initialMsg.IP)

	for {
		// Estrutura para receber as mensagens
		var msg struct {
			Action   string   `json:"action"`
			IP       string   `json:"ip"`
			URLs     []string `json:"urls,omitempty"`
			Interval int      `json:"interval,omitempty"`
			IsPBI    bool     `json:"isPBI"`
		}

		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("❌ Erro ao ler mensagem: %v", err)
			}
			break
		}

		// Log detalhado da mensagem recebida
		log.Printf("📩 Mensagem recebida:")
		log.Printf("   - Ação: %s", msg.Action)
		log.Printf("   - IP: %s", msg.IP)
		if len(msg.URLs) > 0 {
			log.Printf("   - URLs:")
			for i, url := range msg.URLs {
				log.Printf("     %d. %s", i+1, url)
			}
		} else {
			log.Printf("⚠️ Nenhuma URL recebida.")
		}
		if msg.Interval > 0 {
			log.Printf("   - Intervalo: %d segundos", msg.Interval)
		}

		// Processa a mensagem baseado na ação
		switch msg.Action {
		case "display":
			if targetConn, exists := w.clients[msg.IP]; exists {
				// Processa as URLs para obter as screenshots
				var processedURLs []string
				for _, url := range msg.URLs {
					log.Printf("📸 Capturando screenshot de: %s", url)

					// Chame seu ScreenshotService aqui
					if msg.IsPBI {
						screenshotData, err := w.screenshotService.CaptureScreenshotServicePBI(models.RequestBody{Url: url})
						if err != nil {
							log.Printf("❌ Erro ao capturar screenshot de %s: %v", url, err)
							continue
						}
						processedURLs = append(processedURLs, string(screenshotData)) // Converte []byte para string
					} else {
						screenshotData, err := w.screenshotService.CaptureScreenshotServiceGeneric(models.RequestBody{Url: url})
						if err != nil {
							log.Printf("❌ Erro ao capturar screenshot de %s: %v", url, err)
							continue
						}
						processedURLs = append(processedURLs, string(screenshotData)) // Converte []byte para string
					}

					log.Printf("✅ Screenshot capturada com sucesso!")
				}

				// Envia a mensagem com as imagens codificadas em Base64
				displayMsg := struct {
					Action   string   `json:"action"`
					IP       string   `json:"ip"`
					Base64   []string `json:"base64"` // Envia o Base64 da imagem
					Interval int      `json:"interval"`
				}{
					Action:   "display_response",
					IP:       msg.IP,
					Base64:   processedURLs,
					Interval: msg.Interval,
				}

				err := targetConn.WriteJSON(displayMsg)
				if err != nil {
					log.Printf("❌ Erro ao enviar mensagem para %s: %v", msg.IP, err)
				} else {
					log.Printf("✅ Mensagem enviada com sucesso para %s", msg.IP)
				}
			}

		default:
			log.Printf("⚠️ Ação desconhecida: %s", msg.Action)
		}
	}

	// Quando a conexão for fechada, remover do mapa de clients
	for ip, c := range w.clients {
		if c == conn {
			delete(w.clients, ip)
			log.Printf("👋 Cliente desconectado: %s:%s", ip, RASPBERRY_PORT)
			break
		}
	}
}

// Função auxiliar para salvar a screenshot e retornar a URL
// func saveScreenshotAndGetURL(screenshot []byte, originalURL string) string {
// ... código existente ...
// }
