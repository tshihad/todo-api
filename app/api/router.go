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
		r.Get("/todolist", a.handleGetTodoList)
		r.Post("/todolist", a.handlePostTodoList)
		r.Put("/todolist", a.handlePutTodoList)
		r.Delete("/todolist", a.handleDeleteTodoList)
	})
	return r
}
