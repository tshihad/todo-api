package data

import (
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

// NewRepo new repo instance
func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}
