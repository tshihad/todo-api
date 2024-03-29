package api

import (
	"encoding/json"
	"net/http"
	"todo/app/models"

	"github.com/go-chi/chi"
)

func (a *App) handlePostUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	var userid int
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		a.Error(err)
		a.Fail(w, "invalid json format", http.StatusBadRequest)
		return
	}
	userid, err := a.InsertUser(user)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	user.ID = userid
	a.Success(w, http.StatusCreated, user)
}

func (a *App) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	user, err := a.GetUser(userName)
	if err != nil {
		a.Error(err)
		a.Fail(w, "Record not found", http.StatusNotFound)
		return
	}
	a.Success(w, http.StatusOK, user)
}
