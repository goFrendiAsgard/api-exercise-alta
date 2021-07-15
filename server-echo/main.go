package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// go mod init <url>/<package>
// go mod init github.com/<user>/<repo>

func main() {
	e := echo.New()
	e.POST("/", postHandler)
	e.GET("/", getHandler)
	e.Start(":8080")
}

func postHandler(c echo.Context) error {
	c.String(200, "hello world from echo, method: POST")
	return nil
}

func getHandler(c echo.Context) error {
	//c.String(200, "hello world from echo")
	c.JSON(http.StatusOK, map[string]string{
		"name":  "john snow",
		"house": "stark",
	})
	return nil
}
