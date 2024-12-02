package dbmodel

import (
	"gorm.io/gorm"
)

type VisitEntry struct {
	gorm.Model
	Date  string `json:"visit_date"`
	Veto  string `json:"visit_veto"`
	Motif string `json:"visit_motif"`
	IdCat int    `json:"visit_id_cat"`
}

type VisitEntryRepository interface {
	Create(entry *VisitEntry) (*VisitEntry, error)
	FindAll() ([]*VisitEntry, error)
}

type visitEntryRepository struct {
	db *gorm.DB
}

func NewVisitEntryRepository(db *gorm.DB) VisitEntryRepository {
	return &visitEntryRepository{db: db}
}

func (r *visitEntryRepository) Create(entry *VisitEntry) (*VisitEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *visitEntryRepository) FindAll() ([]*VisitEntry, error) {
	var entries []*VisitEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
