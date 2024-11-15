package main

import (
	"log"
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/pkg/cat"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
    router := chi.NewRouter()
    router.Mount("/api/v1/cat", cat.Routes(configuration))
    return router
}

func main() {
    // Initialisation de la configuration
    configuration, err := config.New()
    if err != nil {
        log.Panicln("Configuration error:", err)
    }

    // Initialisation des routes
    router := Routes(configuration)

    log.Println("Serving on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}