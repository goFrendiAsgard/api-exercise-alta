package routes

import (
	"alta/training/controllers"
	"alta/training/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.GET("/users", controllers.GetUserControllers)
	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
	eAuth.GET("/rahasia", controllers.RahasiaController)
	// contoh closure
	nama := "Fawwaz"
	e.GET("/hello", controllers.CreateHelloController(nama))
	e.GET("/hi", controllers.CreateHelloController(nama))
}
