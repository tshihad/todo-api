package data

import (
	"todo/app/models"
)

// Repo repository layer
type Repo interface {
	UserRepo
	TodoListRepo
	TodoRepo
}

// UserRepo user repo
type UserRepo interface {
	InsertUser(models.Users) (int, error)
	GetUser(userName string) (models.Users, error)
}

// TodoListRepo for todolist
type TodoListRepo interface {
	GetTodoList(userID int) ([]models.TodoList, error)
	InsertTodoList(todoList models.TodoList) (models.TodoList, error)
	UpdateTodoList(int, string) error
	DeleteTodoList(todoListID int) error
}

// TodoRepo todo repo
type TodoRepo interface {
	InsertTodo(models.Todo) (models.Todo, error)
	GetTodo(todoListID int, status int) ([]models.Todo, error)
	UpdateTodo(id int, todo models.Todo) error
	DeleteTodo(int) error
}
