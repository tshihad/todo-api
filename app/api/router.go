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

		r.Get("/todolist/{user_id}", a.handleGetTodoList)
		r.Post("/todolist/{user_id}", a.handlePostTodoList)
		r.Put("/todolist/{todo_list_id}", a.handlePutTodoList)
		r.Delete("/todolist/{todo_list_id}", a.handleDeleteTodoList)

		r.Get("/todo/{todo_list_id}/{status:(all|pending|done)}", a.handleGetTodo)
		r.Post("/todo/{todo_list_id}", a.handlePostTodo)
		r.Put("/todo/{todo_id}", a.handlePutTodo)
		r.Delete("/todo/{todo_id}", a.handleDeleteTodo)
	})
	return r
}
