package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, 世界")
	})
	e.Start(":8080")
}
