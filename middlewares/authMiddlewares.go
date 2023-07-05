package middlewares

import "github.com/labstack/echo"

func BasicAuth(username, password string, c *echo.Echo) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
