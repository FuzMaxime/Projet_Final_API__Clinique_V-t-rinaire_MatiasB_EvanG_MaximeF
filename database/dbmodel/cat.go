package dbmodel

import (
	"gorm.io/gorm"
)

type CatEntry struct {
    gorm.Model
    Name string `json:"cat_name"`
    Age   int `json:"cat_age"`
	Race  string `json:"cat_race"`
	Weight int `json:"cat_weight"`
}

type CatEntryRepository interface {
    Create(entry *CatEntry) (*CatEntry, error)
    FindAll() ([]*CatEntry, error)
}

type catEntryRepository struct {
    db *gorm.DB
}

func NewCatEntryRepository(db *gorm.DB) CatEntryRepository {
    return &catEntryRepository{db: db}
}

func (r *catEntryRepository) Create(entry *CatEntry) (*CatEntry, error) {
    if err := r.db.Create(entry).Error; err != nil {
        return nil, err
    }
    return entry, nil
}

func (r *catEntryRepository) FindAll() ([]*CatEntry, error) {
    var entries []*CatEntry
    if err := r.db.Find(&entries).Error; err != nil {
        return nil, err
    }
    return entries, nil
}