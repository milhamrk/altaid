package config

import (
	"alta.id/go-skeleton/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:root@tcp(mamphost:3306)/jekom?charset=utf8&parseTime=True&loc=Local"

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitialMigration()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Buku{})
	DB.AutoMigrate(&models.User{})
}
