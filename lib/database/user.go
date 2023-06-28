package database

import (
	"alta.id/go-skeleton/config"
	"alta.id/go-skeleton/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
