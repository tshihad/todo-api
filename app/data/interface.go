package data

import (
	"todo/app/models"
)

// Repo repository layer
type Repo interface {
	UserRepo
	TodoListRepo
}

// UserRepo user repo
type UserRepo interface {
	InsertUser(models.Users) (int, error)
	GetUser(userName string) (models.Users, error)
}

// TodoListRepo for todolist
type TodoListRepo interface {
	GetTodoLists(userID int) ([]models.TodoList, error)
	InsertTodoList(todoList models.TodoList) (models.TodoList, error)
	UpdateTodoList(name string) error
	DeleteTodoList(todoListID int) error
	InsertTodo(models.Todo) (models.Todo, error)
}
