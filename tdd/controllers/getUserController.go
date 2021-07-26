package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string
	Age  int
}

func GetUserController(c echo.Context) error {
	users := []User{
		{Name: "Anton", Age: 20},
		{Name: "Budi", Age: 10},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}
