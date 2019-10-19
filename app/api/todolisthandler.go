package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/app/models"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func (a *App) handlePostTodoList(w http.ResponseWriter, r *http.Request) {
	var todoList models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&todoList); err != nil {
		a.Error(err)
		a.Fail(w, "Bad request", http.StatusBadRequest)
		return
	}
	todoList, err := a.InsertTodoList(todoList)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to create todolist", http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusCreated, todoList)
}

func (a *App) handleGetTodoList(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid user id format", http.StatusBadRequest)
		return
	}
	todoList, err := a.GetTodoLists(userID)
	if err != nil {
		a.Error(err)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusOK)
			return
		}
		a.Fail(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	a.Success(w, http.StatusOK, todoList)
}

func (a *App) handlePutTodoList(w http.ResponseWriter, r *http.Request) {
	var todoListReq models.TodoListPutReq
	if err := json.NewDecoder(r.Body).Decode(&todoListReq); err != nil {
		a.Error(err)
		a.Fail(w, "Bad request", http.StatusBadRequest)
		return
	}
	if err := a.UpdateTodoList(todoListReq.Name); err != nil {
		if err != nil {
			a.Error(err)
			a.Fail(w, "Failed to update todolist", http.StatusInternalServerError)
			return
		}
	}

}

func (a *App) handleDeleteTodoList(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todoListID, err := strconv.Atoi(id)
	if err != nil {
		a.Error(err)
		a.Fail(w, "invalid todolist id format", http.StatusBadRequest)
		return
	}
	if err := a.DeleteTodoList(todoListID); err != nil {
		a.Error(err)
		a.Fail(w, "Failed to delete todolist", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
