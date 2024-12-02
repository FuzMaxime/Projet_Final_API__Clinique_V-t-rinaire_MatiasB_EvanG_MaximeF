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
	router.Get("/all-treatments", treatmentConfig.GetAllTreatmentHandler)
	router.Get("/one-treatment/{id}", treatmentConfig.GetOneTreatmentHandler)
	router.Put("/update-treatment/{id}", treatmentConfig.UpdateTreatmentHandler)
	router.Delete("/delete-treatment/{id}", treatmentConfig.DeleteTreatmentHandler)

	return router
}
