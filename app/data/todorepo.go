package data

import (
	"todo/app/models"
)

// InsertTodo insert todo to todolist
func (r *RepoImp) InsertTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	return todo, err
}

// GetTodo get  todo from todolist
func (r *RepoImp) GetTodo(todoListID int, status int) ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Where(models.Todo{
		TodoListID: todoListID,
		Status:     status,
	}).Find(&todos).Error
	return todos, err
}

// UpdateTodo updates todo with/without status
func (r *RepoImp) UpdateTodo(id int, todo models.Todo) error {
	return r.db.Model(todo).Where(models.Todo{ID: id}).Update(&todo).Error
}

// DeleteTodo delete todo
func (r *RepoImp) DeleteTodo(id int) error {
	return r.db.Model(models.Todo{}).Delete(models.Todo{ID: id}).Error
}
