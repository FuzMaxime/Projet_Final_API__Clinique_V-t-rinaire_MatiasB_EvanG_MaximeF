package cat

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CatConfig {
	return &CatConfig{configuration}
}

func (config *CatConfig) CreateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat creation request loaded"})
		return
	}

	catEntry := &dbmodel.CatEntry{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	config.CatEntryRepository.Create(catEntry)

	res := &model.CatResponse{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	render.JSON(w, r, res)
}

func (config *CatConfig) GetAllCatsHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *CatConfig) GetOneCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, _ := strconv.Atoi(catId)
	var catTarget *dbmodel.CatEntry

	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			catTarget = cat
		}
	}

	if catTarget == nil {
		render.JSON(w, r, map[string]string{"error": "Cat not found"})
		return
	}
	render.JSON(w, r, catTarget)
}

func (config *CatConfig) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")
	intcatId, err := strconv.Atoi(catId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID"})
		return
	}

	catEntry, err := config.CatEntryRepository.FindByID(uint(intcatId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Cat not found"})
		return
	}

	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat update request loaded"})
		return
	}

	catEntry.Name = req.Name
	catEntry.Age = req.Age
	catEntry.Race = req.Race
	catEntry.Weight = req.Weight

	updatedCat, err := config.CatEntryRepository.Update(catEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update cat"})
		return
	}

	render.JSON(w, r, updatedCat)
}

func (config *CatConfig) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, err := strconv.Atoi(catId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID conversion"})
		return
	}
	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			config.CatEntryRepository.Delete(cat)
			render.JSON(w, r, "Oups, we have kill your cat!")
		}
	}

	render.JSON(w, r, "Cat not found")
}
