package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o router usando Wire
	router, err := InitializeApplication()
	if err != nil {
		log.Fatalf("❌ Falha na inicialização: %v", err)
	}
	// Configurações do Gin
	gin.SetMode(gin.ReleaseMode) // Altere para gin.DebugMode em desenvolvimento

	// Inicia o servidor
	log.Println("🚀 Servidor iniciado na porta 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("💥 Erro ao iniciar servidor: %v", err)
	}
}
