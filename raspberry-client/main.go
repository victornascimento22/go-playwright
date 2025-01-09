package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

type ImagePayload struct {
	Image          string `json:"image"`
	Index          int    `json:"index"`
	TransitionTime int    `json:"transition_time"`
}

const (
	PORT         = "8081"
	DISPLAY_PATH = "/home/loadt/raspberryclient/pi"
)

var (
	fehCmd         *exec.Cmd
	fehMutex       sync.Mutex
	currentIndex   int
	imageCount     int
	transitionTime int
	stopTransition chan bool
)

func init() {
	os.MkdirAll(DISPLAY_PATH, 0755)
	stopTransition = make(chan bool)
}

func startTransitions() {
	log.Printf("🔄 Iniciando transições com intervalo de %d segundos", transitionTime)

	// Para transição anterior se existir
	select {
	case stopTransition <- true:
	default:
	}
	stopTransition = make(chan bool)

	go func() {
		for {
			select {
			case <-stopTransition:
				log.Printf("🛑 Transições interrompidas")
				return
			default:
				showImage(currentIndex)
				log.Printf("🔄 Próxima imagem em %d segundos", transitionTime)
				currentIndex = (currentIndex + 1) % imageCount
				time.Sleep(time.Duration(transitionTime) * time.Second)
			}
		}
	}()
}

func showImage(index int) {
	fehMutex.Lock()
	defer fehMutex.Unlock()

	// Mata processo anterior
	if fehCmd != nil && fehCmd.Process != nil {
		fehCmd.Process.Kill()
		fehCmd.Wait() // Espera o processo terminar
	}

	imagePath := filepath.Join(DISPLAY_PATH, fmt.Sprintf("image_%d.png", index))

	// Verifica se o arquivo existe
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Printf("❌ Imagem não encontrada: %s", imagePath)
		return
	}

	// Executa o feh com mais opções
	fehCmd = exec.Command("feh",
		"-F",                      // Tela cheia
		"--hide-pointer",          // Esconde o cursor
		"--auto-zoom",             // Ajusta zoom automaticamente
		"--force-aliasing",        // Força melhor qualidade
		"--quiet",                 // Reduz logs
		"--on-last-slide", "hold", // Mantém última imagem
		imagePath,
	)

	// Configura ambiente
	fehCmd.Env = append(os.Environ(), "DISPLAY=:0")

	// Captura saída de erro
	var stderr bytes.Buffer
	fehCmd.Stderr = &stderr

	if err := fehCmd.Start(); err != nil {
		log.Printf("❌ Erro ao exibir imagem %d: %v\nErro: %s\n", index, err, stderr.String())
	} else {
		log.Printf("✨ Exibindo imagem %d: %s\n", index, imagePath)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("📥 Recebendo requisição de %s", r.RemoteAddr)

	var payload ImagePayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Printf("❌ Erro ao decodificar JSON: %v", err)
		http.Error(w, "Erro ao ler payload", http.StatusBadRequest)
		return
	}

	log.Printf("✅ Recebido payload: índice=%d, tempo=%ds", payload.Index, payload.TransitionTime)

	// Decodifica a imagem de base64
	imageBytes, err := base64.StdEncoding.DecodeString(payload.Image)
	if err != nil {
		log.Printf("❌ Erro ao decodificar imagem: %v", err)
		http.Error(w, "Erro ao decodificar imagem", http.StatusBadRequest)
		return
	}

	// Salva a imagem
	imagePath := filepath.Join(DISPLAY_PATH, fmt.Sprintf("image_%d.png", payload.Index))
	if err := os.WriteFile(imagePath, imageBytes, 0644); err != nil {
		log.Printf("❌ Erro ao salvar imagem: %v", err)
		http.Error(w, "Erro ao salvar imagem", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Imagem salva em: %s", imagePath)

	// Atualiza configurações
	imageCount = payload.Index + 1
	transitionTime = payload.TransitionTime

	// Inicia transições se for a primeira imagem
	if payload.Index == 0 {
		currentIndex = 0
		startTransitions()
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Imagem recebida e processada com sucesso")
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/ping", handlePing)
	log.Printf("🚀 Webhook rodando em http://localhost:%s\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
