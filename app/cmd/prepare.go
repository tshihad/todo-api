package cmd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func mustPrepareDB() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo password=pass123 sslmode=disable")
}
