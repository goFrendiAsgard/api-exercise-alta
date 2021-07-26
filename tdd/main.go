package main

import (
	"alta/tdd/controllers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.Start(":8080")
}
