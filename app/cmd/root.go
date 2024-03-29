package cmd

import (
	"net/http"
	"sync"

	"github.com/jinzhu/gorm"

	"todo/app/api"
)

// Execute execution of start here
func Execute() {
	db, err := mustPrepareDB()
	if err != nil {
		panic(err)
	}
	serveApp(db)
}

func serveApp(db *gorm.DB) {
	var wg sync.WaitGroup
	app := api.NewApp(db)
	wg.Add(1)
	go func() {
		err := http.ListenAndServe("0.0.0.0:8080", app.Router())
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
