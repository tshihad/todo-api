package models

// Users user
type Users struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:users_name"`
	Email string `json:"email" gorm:"column:email"`
}
