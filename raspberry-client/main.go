package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Screenshot struct {
	ImageData string `json:"imageData"` // Base64 da imagem
	Url       string `json:"url"`       // URL original
}

type Message struct {
	Action      string       `json:"action"`
	IP          string       `json:"ip"`
	Screenshots []Screenshot `json:"screenshots,omitempty"`
	Interval    int          `json:"interval,omitempty"`
	URLs        []string     `json:"urls"` // Removido "omitempty" para garantir que o campo nunca seja ignorado
}

type Response struct {
	Action  string `json:"action"`
	Status  string `json:"status"`
	Message string `json:"message"`
	IP      string `json:"ip"`
}

func main() {
	// Configurar log para incluir arquivo e número da linha
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Conecta ao servidor WebSocket na porta 8080
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/connect"}
	log.Printf("🔗 Conectando ao servidor: %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("❌ Erro ao conectar:", err)
	}
	defer c.Close()

	// Registra o Raspberry
	localIP := "127.0.0.1" // Ajuste para o IP do seu Raspberry
	registerMsg := Message{
		Action: "register",
		IP:     localIP,
	}

	err = c.WriteJSON(registerMsg)
	if err != nil {
		log.Fatal("❌ Erro ao registrar:", err)
	}

	log.Printf("✅ Registrado com sucesso usando IP: %s", localIP)

	// Canal para controlar interrupção
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Canal para mensagens recebidas
	done := make(chan struct{})

	// Goroutine para ler mensagens
	go func() {
		defer close(done)
		for {
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("❌ Erro ao ler mensagem:", err)
				return
			}

			// 🔍 Log detalhado do JSON recebido
			log.Printf("📩 JSON recebido: %+v", msg)

			// Processa a ação recebida
			switch msg.Action {
			case "display_response":
				log.Printf("📷 Recebido comando para exibir %d imagens", len(msg.URLs))

				if len(msg.URLs) == 0 {
					log.Println("⚠️ Nenhuma URL recebida!")
				}

				for i, url := range msg.URLs {
					log.Printf("🖼️ Imagem %d: %s", i+1, url)
					// Aqui você pode implementar a lógica para exibir a imagem
				}

			case "disconnect":
				log.Println("🔌 Recebido comando para desconectar")
				return

			default:
				log.Printf("⚠️ Ação desconhecida recebida: %s", msg.Action)
			}
		}
	}()

	// Aguarda sinal de interrupção
	select {
	case <-interrupt:
		log.Println("🛑 Interrupção recebida, fechando conexão...")
		err := c.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		if err != nil {
			log.Printf("❌ Erro ao enviar mensagem de fechamento: %v", err)
		}
	case <-done:
		log.Println("🔌 Conexão fechada pelo servidor")
	}
}
