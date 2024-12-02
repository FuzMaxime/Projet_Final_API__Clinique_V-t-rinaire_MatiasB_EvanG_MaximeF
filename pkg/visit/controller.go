package Visit

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
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	visitEntry := &dbmodel.VisitEntry{Date: req.Date, Veto: req.Veto, Motif: req.Motif, IdCat: req.IdCat}
	config.VisitEntryRepository.Create(visitEntry)

	res := &model.VisitResponse{}
	render.JSON(w, r, res)
}

func (config *VisitConfig) VisitHistoryHandler(w http.ResponseWriter, r *http.Request) {
	idCat := chi.URLParam(r, "id_cat")
	intIdCat, err := strconv.Atoi(idCat)
	entries, err := config.VisitEntryRepository.FindAll()
	for i := 0; i < len(entries); i++ {
		if entries[i].IdCat != intIdCat {
			entries = append(entries[:i], entries[i+1:]...)
		}
	}
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}
