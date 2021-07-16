package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

func main() {
	db, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT id, title, author from books")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var book Book
		err = results.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Title:", book.Title)
		fmt.Println("Author:", book.Author)
	}
}
