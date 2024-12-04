package treatment

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
		render.JSON(w, r, map[string]string{"error": "Invalid treatment creation request loaded"})
		return
	}

	treatmentEntry := &dbmodel.TreatmentEntry{Medoc: req.Medoc, IdVisit: req.IdVisit}
	config.TreatmentEntryRepository.Create(treatmentEntry)

	res := &model.TreatmentResponse{Medoc: req.Medoc, IdVisit: req.IdVisit}
	render.JSON(w, r, res)
}

func (config *TreatmentConfig) TreatmentHistoryHandler(w http.ResponseWriter, r *http.Request) {

	IdVisit := chi.URLParam(r, "id_visit")
	intIdVisit, err := strconv.Atoi(IdVisit)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID conversion"})
		return
	}
	entries, err := config.TreatmentEntryRepository.FindAll()
	newList := []*dbmodel.TreatmentEntry{}
	for _, value := range entries {
		if value.IdVisit == intIdVisit {
			newList = append(newList, value)
		}
	}
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, newList)
}

func (config *TreatmentConfig) GetAllTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.TreatmentEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *TreatmentConfig) GetOneTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	treatmentId := chi.URLParam(r, "id")

	entries, err := config.TreatmentEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intTreatmentId, err := strconv.Atoi(treatmentId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID conversion"})
		return
	}
	var treatmentTarget *dbmodel.TreatmentEntry

	for _, treatment := range entries {
		if treatment.ID == uint(intTreatmentId) {
			treatmentTarget = treatment
		}
	}

	render.JSON(w, r, treatmentTarget)
}

func (config *TreatmentConfig) UpdateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	treatmentId := chi.URLParam(r, "id")
	intTreatmentId, err := strconv.Atoi(treatmentId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID"})
		return
	}

	treatmentEntry, err := config.TreatmentEntryRepository.FindByID(uint(intTreatmentId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Cat not found"})
		return
	}

	req := &model.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment update request loaded"})
		return
	}

	treatmentEntry.Medoc = req.Medoc
	treatmentEntry.IdVisit = req.IdVisit

	updatedTreatment, err := config.TreatmentEntryRepository.Update(treatmentEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update treatment"})
		return
	}

	render.JSON(w, r, updatedTreatment)
}

func (config *TreatmentConfig) DeleteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	treatmentId := chi.URLParam(r, "id")

	entries, err := config.TreatmentEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intTreatmentId, err := strconv.Atoi(treatmentId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid treatment ID conversion"})
		return
	}

	for _, treatment := range entries {
		if treatment.ID == uint(intTreatmentId) {
			config.TreatmentEntryRepository.Delete(treatment)
			render.JSON(w, r, "You suppressed a treatment!")
		}
	}

	render.JSON(w, r, map[string]string{"error": "Treatment not found"})
}
