package cat

import (
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

	"github.com/go-chi/render"
)

type CatConfig struct {
    *config.Config
}

func New(configuration *config.Config) *CatConfig {
    return &CatConfig{configuration}
}

func (config *CatConfig) CatHandler(w http.ResponseWriter, r *http.Request) {
    req := &model.CatRequest{}
    if err := render.Bind(r, req); err != nil {
        render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
        return
    }

    catEntry := &dbmodel.CatEntry{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
    config.CatEntryRepository.Create(catEntry)

    res := &model.CatResponse{CatAge: catEntry.Age}
    render.JSON(w, r, res)
}

func (config *CatConfig) CatHistoryHandler(w http.ResponseWriter, r *http.Request) {
    entries, err := config.CatEntryRepository.FindAll()
    if err != nil {
        render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
        return
    }
    render.JSON(w, r, entries)
}