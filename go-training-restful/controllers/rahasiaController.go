package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func RahasiaController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "rahasia, cuma bisa diakses yang login",
	})
}
