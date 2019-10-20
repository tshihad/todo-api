package data

import (
	"todo/app/models"
)

func (r *repo) InsertTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

func (r *repo) GetTodo(todoListID int, status int) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where(models.Todo{
		TodoListID: todoListID,
		Status:     status,
	}).Find(&todos).Error
	return todos, err
}

func (r *repo) UpdateTodo(id int, todo models.Todo) error {
	return r.db.Model(todo).Where(models.Todo{ID: id}).Update(&todo).Error
}
func (r *repo) DeleteTodo(id int) error {
	return r.db.Model(models.Todo{}).Delete(models.Todo{ID: id}).Error
}
