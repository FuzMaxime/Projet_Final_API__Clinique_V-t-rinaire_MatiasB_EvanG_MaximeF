package model

import (
	"errors"
	"net/http"
)

type TreatmentRequest struct {
	Medoc   string `json:"treatment_medoc"`
	IdVisit int    `json:"treatment_id_visit"`
}

func (a *TreatmentRequest) Bind(r *http.Request) error {
	if a.IdVisit < 0 {
		return errors.New("treatment_id_visit n'est  pas bien rempli")
	}
	if a.Medoc == "" {
		return errors.New("treatment_medoc n'est  pas bien rempli")
	}
	return nil
}

type TreatmentResponse struct {
}
