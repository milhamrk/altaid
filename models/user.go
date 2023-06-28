package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
