package database

import (
	"log"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.CatEntry{},
		&dbmodel.TreatmentEntry{},
		&dbmodel.VisitEntry{},
	)
	log.Println("Database migrated successfully")
}
