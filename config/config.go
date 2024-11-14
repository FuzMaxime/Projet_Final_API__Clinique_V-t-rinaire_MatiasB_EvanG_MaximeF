package config

import (
	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
    // Connexion aux repositories
    CatEntryRepository     dbmodel.CatEntryRepository
}

func New() (*Config, error) {
    config := Config{}

    // Initialisation de la connexion à la base de données
    databaseSession, err := gorm.Open(sqlite.Open("vet-clinic-api.db"), &gorm.Config{})
    if err != nil {
        return &config, err
    }

    // Migration des modèles
    database.Migrate(databaseSession)

    // Initialisation des repositories
    config.CatEntryRepository = dbmodel.NewCatEntryRepository(databaseSession)

    return &config, nil
}