package core

import (
	"fmt"
	"log"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// NewSQLMock is mock db
func NewSQLMock() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	gormDB.SingularTable(true)

	return mock, gormDB.Set("gorm:update_column", true)
}

// FixedFullRe return fix full query
func FixedFullRe(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}
