package data

import (
	"github.com/jinzhu/gorm"
)

// RepoImp implements repo instance
type RepoImp struct {
	db *gorm.DB
}

// NewRepo new repo instance
func NewRepo(db *gorm.DB) *RepoImp {
	return &RepoImp{
		db: db,
	}
}
