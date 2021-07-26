package controller

import (
	"alta/tddmock/model"

	"github.com/labstack/echo/v4"
)

func CreateBookController(bookModel model.BookModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: lengkapi sendiri
		return nil
	}
}
