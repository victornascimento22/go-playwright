package infra

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketManager gerencia conexões WebSocket
type WebSocketManager struct {
	connections map[string]*websocket.Conn // RaspberryIP -> Conexão
	mu          sync.RWMutex               // Para acesso seguro ao map
}

// Cria novo manager
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string]*websocket.Conn),
	}
}

// Adiciona nova conexão
func (wm *WebSocketManager) AddConnection(raspberryIP string, conn *websocket.Conn) {
	wm.mu.Lock()
	wm.connections[raspberryIP] = conn
	wm.mu.Unlock()
	log.Printf("📡 Nova conexão WebSocket: %s", raspberryIP)
}

// Remove conexão
func (wm *WebSocketManager) RemoveConnection(raspberryIP string) {
	wm.mu.Lock()
	delete(wm.connections, raspberryIP)
	wm.mu.Unlock()
	log.Printf("❌ Conexão WebSocket fechada: %s", raspberryIP)
}

// Pega conexão existente
func (wm *WebSocketManager) GetConnection(raspberryIP string) (*websocket.Conn, bool) {
	wm.mu.RLock()
	conn, exists := wm.connections[raspberryIP]
	wm.mu.RUnlock()
	return conn, exists
}
