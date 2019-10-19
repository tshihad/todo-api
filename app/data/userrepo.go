package data

import (
	"todo/app/models"
)

func (r *repo) InsertUser(user models.Users) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return -1, err
	}
	return user.ID, nil
}

func (r *repo) GetUser(userName string) (models.Users, error) {
	var user models.Users
	if err := r.db.Where(models.Users{Name: userName}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
