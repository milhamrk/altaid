package routes

import (
	"alta.id/go-skeleton/controllers"
	"alta.id/go-skeleton/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)

	euth := e.Group("")
	euth.Use(middleware.BasicAuth(middlewares.BasicAuthDB))

	return e
}
