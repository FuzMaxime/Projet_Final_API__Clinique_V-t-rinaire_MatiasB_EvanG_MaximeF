package cat

import (
	"net/http"
	"strconv"
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

func (config *CatConfig) CreateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	catEntry := &dbmodel.CatEntry{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	config.CatEntryRepository.Create(catEntry)

	res := &model.CatResponse{Name: catEntry.Name, Age: catEntry.Age, Race: catEntry.Race, Weight: catEntry.Weight}
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
	catId := r.URL.Query().Get("id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, err := strconv.Atoi(catId)
	var catTarget *dbmodel.CatEntry

	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			catTarget = cat
			return
		}
	}

	render.JSON(w, r, catTarget)
}

func (config *CatConfig) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := r.URL.Query().Get("id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, err := strconv.Atoi(catId)

	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			config.CatEntryRepository.Update(cat)
			return
		}
	}

	render.JSON(w, r, "Your cat is update!")
}

func (config *CatConfig) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := r.URL.Query().Get("id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, err := strconv.Atoi(catId)

	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			config.CatEntryRepository.Delete(cat)
			return
		}
	}

	render.JSON(w, r, "Oups, we have kill your cat!")
}
