package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

func CreateHelloController(name string) echo.HandlerFunc {
	return func(c echo.Context) error {
		greetings := fmt.Sprintf("Hello %s", name)
		return c.String(200, greetings)
	}
}
