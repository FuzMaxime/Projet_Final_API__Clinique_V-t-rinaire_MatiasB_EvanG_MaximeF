package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
    catConfig := New(configuration)
    router := chi.NewRouter()

    router.Post("/age-in-cat-years", catConfig.CatHandler)
    router.Get("/history", catConfig.CatHistoryHandler)

    return router
}