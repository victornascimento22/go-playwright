package models

import "time"

// Display representa uma TV/Display conectada a um Raspberry Pi
type Display struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`         // Nome da TV (ex: "TV Operação 1")
	RaspberryIP string    `json:"raspberry_ip"` // IP do Raspberry
	URL         string    `json:"url"`          // URL atual do dashboard
	LastUpdate  time.Time `json:"last_update"`
	Status      string    `json:"status"` // online/offline/error
}

// RequestDisplay representa a requisição para atualizar um display
type RequestDisplay struct {
	URL string `json:"url" binding:"required"`
}
type URLConfig struct {
	URL    string `json:"url"`
	Source string `json:"source"` // "pbi" ou "generic"
}
type DisplayConfig struct {
	URLs           []URLConfig `json:"urls"`            // Lista de URLs para exibir
	TransitionTime int         `json:"transition_time"` // Tempo em segundos entre transições
	RaspberryIP    string      `json:"raspberry_ip"`    // IP do Raspberry
}
