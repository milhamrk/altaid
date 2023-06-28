package models

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Judul     string `json:"judul" form:"name"`
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit" form:"penerbit"`
}
