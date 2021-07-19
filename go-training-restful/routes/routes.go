package routes

import (
	"alta/training/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserControllers)
	return e
}
