package cmd

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func mustPrepareDB() (*sql.DB, error) {
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo password=pass123")
}
