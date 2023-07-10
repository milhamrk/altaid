package controllers

import (
	"fmt"
	"net/http"

	"alta.id/go-skeleton/lib/database"
	"alta.id/go-skeleton/models"
	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	fmt.Println("HALLO")
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
