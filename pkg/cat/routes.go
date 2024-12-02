package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	catConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/create-cat", catConfig.CreateCatHandler)
	router.Get("/all-cats", catConfig.GetAllCatsHandler)
	router.Get("/one-cat/{id}", catConfig.GetOneCatHandler)
	router.Put("/update-cat/{id}", catConfig.UpdateCatHandler)
	router.Delete("/delete-cat/{id}", catConfig.DeleteCatHandler)

	return router
}
