package data

import (
	"todo/app/models"
)

func (r *repo) InsertTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repo) GetTodos(todoListID int) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where(models.Todo{TodoListID: todoListID}).Find(&todos).Error
	return todos, err
}

func (r *repo) UpdateTodo(id int) (models.Todo, error) {
	var todo models.Todo
	return todo, nil
}
