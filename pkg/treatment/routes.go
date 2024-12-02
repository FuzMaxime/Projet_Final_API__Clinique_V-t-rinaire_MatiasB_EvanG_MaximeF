package treatment

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	treatmentConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/treatment_create", treatmentConfig.TreatmentHandler)
	router.Get("/visit-treatments/{id_visit}", treatmentConfig.TreatmentHistoryHandler)
	router.Get("/all-cats", treatmentConfig.GetAllTreatmentHandler)
	router.Get("/one-cat/{id}", treatmentConfig.GetOneTreatmentHandler)
	router.Put("/update-cat/{id}", treatmentConfig.UpdateTreatmentHandler)
	router.Delete("/delete-cat/{id}", treatmentConfig.DeleteTreatmentHandler)

	return router
}
