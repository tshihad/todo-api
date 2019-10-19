package models

// Users user
type Users struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:users_name"`
	Email string `json:"email" gorm:"column:email"`
}

// TodoList todo list
type TodoList struct {
	ID     int    `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:todo_name"`
	UserID int    `json:"user_id" gorm:"column:user_id"`
}

type TodoListPutReq struct {
	Name string
}

type Todo struct {
	ID         int    `json:"id" gorm:"column:id"`
	Data       string `json:"data" gorm:"column:data"`
	Status     int    `json:"status" gorm:"column:status"`
	TodoListID int    `json:"todo_list_id" gorm:"column:todo_list_id"`
}
