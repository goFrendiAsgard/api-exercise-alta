package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// go mod init <url>/<package>
// go mod init github.com/<user>/<repo>

func main() {
	e := echo.New()
	e.POST("/", postHandler)
	e.GET("/", getHandler)
	e.GET("/error", triggerError)

	e.GET("/users/:userId/order/:orderId", func(c echo.Context) error {
		// curl localhost:8080/users/u001/order/o001
		userId := c.Param("userId")
		orderId := c.Param("orderId")
		return c.JSON(http.StatusOK, map[string]string{
			"user":  userId,
			"order": orderId,
		})
	})

	e.GET("/users", func(c echo.Context) error {
		// curl 'localhost:8080/users?limit=5&offset=7'
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		return c.JSON(http.StatusOK, map[string]string{
			"limit":  limit,
			"offset": offset,
		})
	})

	e.POST("/users", func(c echo.Context) error {
		// curl -X POST --form name=tono --form address=malang 'localhost:8080/users'
		name := c.FormValue("name")
		address := c.FormValue("address")
		return c.JSON(http.StatusOK, map[string]string{
			"name":    name,
			"address": address,
		})
	})

	e.POST("/orders", func(c echo.Context) error {
		// curl -X POST --data-raw '{"qty": "4", "item": "es kopi susu"}' 'localhost:8080/orders'
		order := map[string]string{}
		c.Bind(&order)
		return c.JSON(http.StatusOK, order)
	})

	e.Start(":8080")
}

func triggerError(c echo.Context) error {
	return fmt.Errorf("pokoknya error")
}

func postHandler(c echo.Context) error {
	return c.String(200, "hello world from echo, method: POST")
}

func getHandler(c echo.Context) error {
	//return c.String(200, "hello world from echo")
	return c.JSON(http.StatusOK, map[string]string{
		"name":  "john snow",
		"house": "stark",
	})
}
