package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/app/models"

	"github.com/go-chi/chi"
)

func (a *App) handlePostTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	listID := chi.URLParam(r, "todo_list_id")
	todoListID, err := strconv.Atoi(listID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	if todo.Status < 1 || todo.Status > 2 {
		a.Fail(w, "invalid status", http.StatusBadRequest)
		return
	}
	todo.TodoListID = todoListID
	todo, err = a.InsertTodo(todo)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusCreated, todo)
}

func (a *App) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	listID := chi.URLParam(r, "todo_list_id")
	status := chi.URLParam(r, "status")

	var statusID int
	todoListID, err := strconv.Atoi(listID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	switch status {
	case "all":
		statusID = 0
	case "pending":
		statusID = 1
	case "done":
		statusID = 2
	}
	todo, err := a.GetTodo(todoListID, statusID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to get todo", http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusOK, todo)
}

func (a *App) handlePutTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	listID := chi.URLParam(r, "todo_id")
	todoListID, err := strconv.Atoi(listID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := a.UpdateTodo(todoListID, todo); err != nil {
		a.Error(err)
		a.Fail(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *App) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	listID := chi.URLParam(r, "todo_id")
	todoListID, err := strconv.Atoi(listID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = a.DeleteTodo(todoListID)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to get todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
