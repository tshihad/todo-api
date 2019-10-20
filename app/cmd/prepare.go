package cmd

import (
	"os"

	"github.com/jinzhu/gorm"
	// initialize postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func mustPrepareDB() (*gorm.DB, error) {
	// TODO: must read following configuration from config files/environments
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "db"
	}
	db, err := gorm.Open("postgres", "host="+host+" port=5432 user=postgres dbname=todo password=pass123 sslmode=disable")
	db.SingularTable(true)
	return db, err
}
