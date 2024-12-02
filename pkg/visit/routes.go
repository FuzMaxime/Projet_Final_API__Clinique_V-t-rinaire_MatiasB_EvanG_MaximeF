package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	visitConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/visit_create", visitConfig.VisitHandler)
	router.Get("/history/{id_cat}", visitConfig.VisitHistoryHandler)

	return router
}
