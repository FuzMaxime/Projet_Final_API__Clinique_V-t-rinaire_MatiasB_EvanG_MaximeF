package dbmodel

import (
	"gorm.io/gorm"
)

type TreatmentEntry struct {
	gorm.Model
	Medoc   string `json:"treatment_medoc"`
	IdVisit int    `json:"treatment_id_visit"`
}

type TreatmentEntryRepository interface {
	Create(entry *TreatmentEntry) (*TreatmentEntry, error)
	FindAll() ([]*TreatmentEntry, error)
}

type treatmentEntryRepository struct {
	db *gorm.DB
}

func NewTreatmentEntryRepository(db *gorm.DB) TreatmentEntryRepository {
	return &treatmentEntryRepository{db: db}
}

func (r *treatmentEntryRepository) Create(entry *TreatmentEntry) (*TreatmentEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *treatmentEntryRepository) FindAll() ([]*TreatmentEntry, error) {
	var entries []*TreatmentEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
