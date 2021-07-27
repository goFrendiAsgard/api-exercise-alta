package controller

import (
	"alta/tddmock/model"
	"fmt"

	"github.com/labstack/echo"
)

func CreateGetBookController(bookModel model.BookModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		books := bookModel.Get()
		return c.JSON(200, books)
	}
}

func CreatePostBookController(bookModel model.BookModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		var book model.Book
		fmt.Printf("%#v", book)
		return c.JSON(200, book)
	}
}
