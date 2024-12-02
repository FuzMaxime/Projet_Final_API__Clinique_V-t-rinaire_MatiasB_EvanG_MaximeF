package Treatment

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

	"github.com/go-chi/chi/v5"
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

	treatmentEntry := &dbmodel.TreatmentEntry{Medoc: req.Medoc, IdVisit: req.IdVisit}
	config.TreatmentEntryRepository.Create(treatmentEntry)

	res := &model.TreatmentResponse{}
	render.JSON(w, r, res)
}

func (config *TreatmentConfig) TreatmentHistoryHandler(w http.ResponseWriter, r *http.Request) {

	IdVisit := chi.URLParam(r, "id_visit")
	intIdVisit, err := strconv.Atoi(IdVisit)
	entries, err := config.VisitEntryRepository.FindAll()
	for i := 0; i < len(entries); i++ {
		if entries[i].IdCat != intIdVisit {
			entries = append(entries[:i], entries[i+1:]...)
		}
	}

	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}
