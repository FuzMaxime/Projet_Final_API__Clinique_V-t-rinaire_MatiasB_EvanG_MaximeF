package visit

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitConfig struct {
	*config.Config
}

func New(configuration *config.Config) *VisitConfig {
	return &VisitConfig{configuration}
}

func (config *VisitConfig) VisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit creation request loaded"})
		return
	}

	visitEntry := &dbmodel.VisitEntry{Date: req.Date, Veto: req.Veto, Motif: req.Motif, IdCat: req.IdCat}
	config.VisitEntryRepository.Create(visitEntry)

	res := &model.VisitResponse{Date: req.Date, Veto: req.Veto, Motif: req.Motif, IdCat: req.IdCat}
	render.JSON(w, r, res)
}

func (config *VisitConfig) VisitHistoryHandler(w http.ResponseWriter, r *http.Request) {
	idCat := chi.URLParam(r, "id_cat")
	intIdCat, err := strconv.Atoi(idCat)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID conversion"})
		return
	}
	entries, err := config.VisitEntryRepository.FindAll()
	newList := []*dbmodel.VisitEntry{}
	for _, value := range entries {
		if value.IdCat == intIdCat {
			newList = append(newList, value)
		}
	}
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, newList)
}

func (config *VisitConfig) GetAllVisitHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.VisitEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *VisitConfig) GetOneVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitId := chi.URLParam(r, "id")

	entries, err := config.VisitEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intVisitId, err := strconv.Atoi(visitId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID conversion"})
		return
	}
	var visitTarget *dbmodel.VisitEntry

	for _, visit := range entries {
		if visit.ID == uint(intVisitId) {
			visitTarget = visit
		}
	}

	if visitTarget == nil {
		render.JSON(w, r, map[string]string{"error": "Visit not found"})
		return
	}

	render.JSON(w, r, visitTarget)
}

func (config *VisitConfig) UpdateVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitId := chi.URLParam(r, "id")
	intVisitId, err := strconv.Atoi(visitId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID conversion"})
		return
	}

	visitEntry, err := config.VisitEntryRepository.FindByID(uint(intVisitId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Visit not found"})
		return
	}

	req := &model.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit update request loaded"})
		return
	}

	visitEntry.Date = req.Date
	visitEntry.Motif = req.Motif
	visitEntry.Veto = req.Veto
	visitEntry.IdCat = req.IdCat

	updatedVisit, err := config.VisitEntryRepository.Update(visitEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update visit"})
		return
	}

	render.JSON(w, r, updatedVisit)
}

func (config *VisitConfig) DeleteVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitId := chi.URLParam(r, "id")

	entries, err := config.VisitEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intvisitId, err := strconv.Atoi(visitId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid visit ID conversion"})
		return
	}

	for _, visit := range entries {
		if visit.ID == uint(intvisitId) {
			config.VisitEntryRepository.Delete(visit)
			render.JSON(w, r, "You suppressed a visit!")
			return
		}
	}

	render.JSON(w, r, map[string]string{"error": "Visit not found"})
}
