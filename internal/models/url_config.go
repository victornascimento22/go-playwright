package models

type URLConfig struct {
	URL    string `json:"url"`
	Source string `json:"source"` // "pbi" ou "generic"
}
