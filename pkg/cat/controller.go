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
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
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

	render.JSON(w, r, catTarget)
}
func (config *CatConfig) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'ID depuis les paramètres de l'URL
	catId := chi.URLParam(r, "id")
	intcatId, err := strconv.Atoi(catId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid cat ID"})
		return
	}

	// Validation et binding du payload
	req := &model.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	// Recherche de l'entrée dans le repository par ID
	catEntry := &dbmodel.CatEntry{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	entries, err := config.CatEntryRepository.FindAll()
	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			config.CatEntryRepository.Update(catEntry)
		}
	}

	// Retour de réponse en cas de succès
	render.JSON(w, r, map[string]string{"message": "Update successful"})
}

func (config *CatConfig) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "id")

	entries, err := config.CatEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcatId, err := strconv.Atoi(catId)

	for _, cat := range entries {
		if cat.ID == uint(intcatId) {
			config.CatEntryRepository.Delete(cat)
		}
	}

	render.JSON(w, r, "Oups, we have kill your cat!")
}
