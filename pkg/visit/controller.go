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

	res := &model.VisitResponse{Date: req.Date, Veto: req.Veto, Motif: req.Motif, IdCat: req.IdCat}
	render.JSON(w, r, res)
}

func (config *VisitConfig) VisitHistoryHandler(w http.ResponseWriter, r *http.Request) {
	idCat := chi.URLParam(r, "id_cat")
	intIdCat, _ := strconv.Atoi(idCat)
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
