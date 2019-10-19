package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router chi router
func (a *App) Router() http.Handler {
	r := chi.NewRouter()
	r.Route("/app", func(r chi.Router) {
		r.Post("/user", a.handlePostUser)
		r.Get("/user/{userName}", a.handleGetUser)
	})
	return r
}
