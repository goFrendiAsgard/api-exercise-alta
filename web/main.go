package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name    string
	Address string
}

var users []User

func getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func addUser(c echo.Context) error {
	newUser := User{}
	c.Bind(&newUser)
	users = append(users, newUser)
	return c.JSON(http.StatusOK, newUser)
}

func main() {
	users = make([]User, 0)
	e := echo.New()
	e.Server.ListenAndServe()
	e.GET("users", getUser)
	e.POST("users", addUser)
	e.Start(":3000")
}
