package api

import (
	"os"
	"todo/app/core"
	"todo/app/data"

	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
)

// App main app
type App struct {
	data.Repo
	logrus.FieldLogger
	core.Response
}

// NewApp new app instance
func NewApp(db *gorm.DB) App {
	repo := data.NewRepo(db)
	logger := core.NewLogger(logrus.InfoLevel, os.Stdout)
	return App{
		Repo:        repo,
		FieldLogger: logger,
		Response:    core.Response{logger},
	}
}
