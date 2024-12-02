package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	visitConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/visit_create", visitConfig.VisitHandler)
	router.Get("/cat-visits/{id_cat}", visitConfig.VisitHistoryHandler)
	router.Get("/all-cats", visitConfig.GetAllVisitHandler)
	router.Get("/one-cat/{id}", visitConfig.GetOneVisitHandler)
	router.Put("/update-cat/{id}", visitConfig.UpdateVisitHandler)
	router.Delete("/delete-cat/{id}", visitConfig.DeleteVisitHandler)

	return router
}
