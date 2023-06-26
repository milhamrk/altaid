package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:root@/jekom?charset=utf8&parseTime=True&loc=Local"

type Buku struct {
	gorm.Model
	Judul     string `json:"judul" form:"name"`
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit" form:"penerbit"`
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&Buku{})
}

func GetBooks(context echo.Context) error {
	var buku []Buku
	if err := DB.Find(&buku).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, buku)
}

func CreateBooks(context echo.Context) error {
	var params Buku
	err := context.Bind(&params)
	if err != nil {
		return err
	}

	if err := DB.Create(&params).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusCreated, params)
}

func main() {
	InitDB()
	InitialMigration()

	e := echo.New()

	// Routing
	e.GET("/books", GetBooks)
	e.POST("/create-books", CreateBooks)

	e.Start(":8000")
}
