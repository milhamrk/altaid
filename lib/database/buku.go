package database

import (
	"alta.id/go-skeleton/config"
	"alta.id/go-skeleton/models"
)

func GetBuku() (interface{}, error) {
	var bukus []models.Buku

	if e := config.DB.Find(&bukus).Error; e != nil {
		return nil, e
	}

	return bukus, nil
}
