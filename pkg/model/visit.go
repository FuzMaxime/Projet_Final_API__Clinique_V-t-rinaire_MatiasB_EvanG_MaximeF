package model

import (
	"errors"
	"net/http"
)

type VisitRequest struct {
	Date  string `json:"visit_date"`
	Motif string `json:"visit_motif"`
	Veto  string `json:"visit_veto"`
	IdCat int    `json:"visit_id_cat"`
}

func (a *VisitRequest) Bind(r *http.Request) error {
	if a.Date == "" {
		return errors.New("visit_date must be a string")
	}
	if a.Motif == "" {
		return errors.New("visit_motif must be a string")
	}
	if a.Veto == "" {
		return errors.New("visit_veto must be a string")
	}
	if a.IdCat < 0 {
		return errors.New("visit_id_cat must be a positive number")
	}
	return nil
}

type VisitResponse struct {
	Date  string `json:"visit_date"`
	Motif string `json:"visit_motif"`
	Veto  string `json:"visit_veto"`
	IdCat int    `json:"visit_id_cat"`
}
