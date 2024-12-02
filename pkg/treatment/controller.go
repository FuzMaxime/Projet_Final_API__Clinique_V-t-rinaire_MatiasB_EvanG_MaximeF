package Treatment

import (
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

	"github.com/go-chi/render"
)

type TreatmentConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TreatmentConfig {
	return &TreatmentConfig{configuration}
}

func (config *TreatmentConfig) TreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	TreatmentEntry := &dbmodel.TreatmentEntry{Medoc: req.Medoc, IdVisit: req.IdVisit}
	config.TreatmentEntryRepository.Create(treatmentEntry)

	res := &model.TreatmentResponse{TreatmentAge: TreatmentEntry.Age}
	render.JSON(w, r, res)
}

func (config *TreatmentConfig) TreatmentHistoryHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.TreatmentEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}
