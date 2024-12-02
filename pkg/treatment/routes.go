package Treatment

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	treatmentConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/treatment_create", treatmentConfig.TreatmentHandler)
	router.Get("/history/{id_visit}", treatmentConfig.TreatmentHistoryHandler)

	return router
}
