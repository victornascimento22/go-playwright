package models

import "time"

type Aniversariantes struct {
	Nome_completo              string    `json:"nome_completo"`
	Nome_cracha                string    `json:"nome_cracha"`
	Aniversario_vida           time.Time `json:"aniversario_vida"`
	Aniversario_empresa        time.Time `json:"aniversario_empresa"`
	Url_aniversario_vida_tv    string    `json:"url_aniversario_vida_tv"`
	Url_aniversario_empresa_tv string    `json:"url_aniversario_empresa_tv"`
}
