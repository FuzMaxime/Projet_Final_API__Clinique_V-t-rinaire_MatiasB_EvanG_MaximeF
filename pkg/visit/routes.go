package visit

import (
	"vet-clinic-api/config"
	"vet-clinic-api/pkg/treatment"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	visitConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/visit_create", visitConfig.VisitHandler)
	router.Get("/cat-visits/{id_cat}", visitConfig.VisitHistoryHandler)
	router.Get("/all-visits", visitConfig.GetAllVisitHandler)
	router.Get("/one-visit/{id}", visitConfig.GetOneVisitHandler)
	router.Put("/update-visit/{id}", visitConfig.UpdateVisitHandler)
	router.Delete("/delete-visit/{id}", visitConfig.DeleteVisitHandler)
	
	router.Get("/one-visit/{id_visit}/treatments", treatment.New(configuration).TreatmentHistoryHandler)

	return router
}
