package data

import (
	"todo/app/models"
)

func (r *repo) InsertTodoList(todoList models.TodoList) (models.TodoList, error) {
	err := r.db.Create(&todoList).Error
	return todoList, err
}

func (r *repo) GetTodoList(userID int) ([]models.TodoList, error) {
	todoList := []models.TodoList{}
	err := r.db.Where(models.TodoList{UserID: userID}).Find(&todoList).Error
	return todoList, err
}

func (r *repo) UpdateTodoList(todoListID int, name string) error {
	return r.db.Model(models.TodoList{}).Where(models.TodoList{ID: todoListID}).Update(models.TodoList{Name: name}).Error
}

func (r *repo) DeleteTodoList(todoListID int) error {
	return r.db.Model(models.TodoList{}).Delete(models.TodoList{ID: todoListID}).Error
}