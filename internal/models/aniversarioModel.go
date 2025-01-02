package models

import "time"

// Aniversariantes represents birthday information for employees, including their life and work anniversaries.
// Fields:
// - Nome_completo: Full name of the employee.
// - Nome_cracha: Name displayed on the employee's badge.
// - Aniversario_vida: Employee's life birthday (date of birth).
// - Aniversario_empresa: Work anniversary date, stored as a string for flexible input handling.
// - Email: Employee's email address.
// - Url_aniversario_vida_tv: URL for displaying life birthday information on TV or screens.
// - Url_aniversario_empresa_tv: URL for displaying work anniversary information on TV or screens.
type Aniversariantes struct {
	Nome_completo              string    `json:"nome_completo"`              // Full name of the employee.
	Nome_cracha                string    `json:"nome_cracha"`                // Badge name of the employee.
	Aniversario_vida           time.Time `json:"aniversario_vida"`           // Life birthday of the employee.
	Aniversario_empresa        string    `json:"aniversario_empresa"`        // Work anniversary date as a string.
	Email                      string    `json:"email"`                      // Employee's email.
	Url_aniversario_vida_tv    string    `json:"url_aniversario_vida_tv"`    // URL for life birthday display.
	Url_aniversario_empresa_tv string    `json:"url_aniversario_empresa_tv"` // URL for work anniversary display.
}

// FormatDate converts the Aniversario_empresa field, which is stored as a string,
// into a formatted date in the "DD/MM/YYYY" format.
// Returns:
// - A string representing the formatted date or an empty string if parsing fails.
func (x *Aniversariantes) FormatDate() string {
	// Parse the Aniversario_empresa field as an ISO 8601 date-time string.
	t, _ := time.Parse("2006-01-02T15:04:05Z", x.Aniversario_empresa)

	// Format the parsed date into "DD/MM/YYYY".
	return t.Format("02/01/2006")
}
