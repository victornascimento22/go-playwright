package models

import "time"

type Aniversariantes struct {
	Nome_completo              string    `json:"nome_completo"`
	Nome_cracha                string    `json:"nome_cracha"`
	Aniversario_vida           time.Time `json:"aniversario_vida"`
	Aniversario_empresa        string    `json:"aniversario_empresa"`
	Email                      string    `json:"email"`
	Url_aniversario_vida_tv    string    `json:"url_aniversario_vida_tv"`
	Url_aniversario_empresa_tv string    `json:"url_aniversario_empresa_tv"`
}

func (x *Aniversariantes) FormatDate() string {
	t, _ := time.Parse("2006-01-02T15:04:05Z", x.Aniversario_empresa)
	return t.Format("02/01/2006")
}
