package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Buku struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul" validation:"numeric,required"`
	Pengarang string `json:"pengarang"`
}

var perpustakan []Buku

func GetBooks(context echo.Context) error {
	return context.JSON(http.StatusOK, perpustakan)
}

func CreateBooks(context echo.Context) error {
	var params Buku
	err := context.Bind(&params)
	if err != nil {
		return err
	}

	perpustakan = append(perpustakan, params)
	return context.JSON(http.StatusCreated, perpustakan)
}

func GetBookDetail(context echo.Context) error {
	id := context.Param("id")
	for _, buku := range perpustakan {
		if strconv.Itoa(buku.ID) == id {
			return context.JSON(http.StatusOK, buku)
		}
	}
	return context.JSON(http.StatusNotFound, "Not Found")
}

// Make delete function

func main() {
	e := echo.New()
	perpustakan = []Buku{
		{ID: 1, Judul: "Halo Book", Pengarang: "Siti"},
	}

	// Routing
	e.GET("/books", GetBooks)
	e.GET("/books/:id", GetBookDetail)
	e.POST("/create-books", CreateBooks)

	e.Start(":8000")
}
