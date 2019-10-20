package data

import (
	"todo/app/models"
)

// InsertTodoList insert todo list
func (r *RepoImp) InsertTodoList(todoList models.TodoList) (models.TodoList, error) {
	err := r.db.Create(&todoList).Error
	return todoList, err
}

// GetTodoList get todo list
func (r *RepoImp) GetTodoList(userID int) ([]models.TodoList, error) {
	todoList := []models.TodoList{}
	err := r.db.Where(models.TodoList{UserID: userID}).Find(&todoList).Error
	return todoList, err
}

// UpdateTodoList updates todo list
func (r *RepoImp) UpdateTodoList(todoListID int, name string) error {
	return r.db.Model(models.TodoList{}).Where(models.TodoList{ID: todoListID}).Update(models.TodoList{Name: name}).Error
}

// DeleteTodoList delete todolist
func (r *RepoImp) DeleteTodoList(todoListID int) error {
	return r.db.Model(models.TodoList{}).Delete(models.TodoList{ID: todoListID}).Error
}
