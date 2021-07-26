package model

import "gorm.io/gorm"

type Book struct {
	gorm.DB
	Title string
}

type BookModel interface {
	Get() []Book
	Insert(Book) error
}

// ==== model utk manipulasi db

type GormBookModel struct {
	db *gorm.DB
}

func (m *GormBookModel) Get() []Book {
	var books []Book
	m.db.Find(&books)
	return books
}

func (m *GormBookModel) Insert(book Book) error {
	tx := m.db.Save(&book)
	return tx.Error
}

func NewGormBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}

// ==== model utk mock testing

type MockBookModel struct {
	books []Book
}

func (m *MockBookModel) Get() []Book {
	return m.books
}

func (m *MockBookModel) Insert(book Book) error {
	m.books = append(m.books, book)
	return nil
}

func NewMockBookModel() *MockBookModel {
	return &MockBookModel{
		books: []Book{},
	}
}
