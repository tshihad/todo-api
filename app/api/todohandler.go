package api

import (
	"encoding/json"
	"net/http"
	"todo/app/models"
)

func (a *App) handlePostTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	todo, err := a.InsertTodo(todo)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusCreated, todo)
}

func (a *App) handleGetTodo(w http.ResponseWriter, r *http.Request) {

}

func (a *App) handlePutTodo(w http.ResponseWriter, r *http.Request) {

}

func (a *App) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {

}
