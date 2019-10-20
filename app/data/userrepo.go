package data

import (
	"todo/app/models"
)

// InsertUser create new user
func (r *RepoImp) InsertUser(user models.Users) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return -1, err
	}
	return user.ID, nil
}

// GetUser get user
func (r *RepoImp) GetUser(userName string) (models.Users, error) {
	var user models.Users
	if err := r.db.Where(models.Users{Name: userName}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
