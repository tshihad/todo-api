package data

import (
	"todo/app/models"
)

// Repo repository layer
type Repo interface {
	UserRepo
}

// UserRepo user repo
type UserRepo interface {
	InsertUser(models.Users) (int, error)
	GetUser(userName string) (models.Users, error)
}
