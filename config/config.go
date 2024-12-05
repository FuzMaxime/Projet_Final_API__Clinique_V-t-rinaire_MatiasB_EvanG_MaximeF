package config

import (
	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	// Connexion aux repositories
	CatEntryRepository       dbmodel.CatEntryRepository
	TreatmentEntryRepository dbmodel.TreatmentEntryRepository
	VisitEntryRepository     dbmodel.VisitEntryRepository
}

func New() (*Config, error) {
	config := Config{}

	// Initialisation de la connexion à la base de données
	databaseSession, err := gorm.Open(sqlite.Open("database/vet-clinic-api.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories
	config.CatEntryRepository = dbmodel.NewCatEntryRepository(databaseSession)
	config.TreatmentEntryRepository = dbmodel.NewTreatmentEntryRepository(databaseSession)
	config.VisitEntryRepository = dbmodel.NewVisitEntryRepository(databaseSession)

	return &config, nil
}
