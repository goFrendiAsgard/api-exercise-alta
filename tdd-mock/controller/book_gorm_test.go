package controller

import (
	"alta/tddmock/config"
	"alta/tddmock/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createTestDB() (*gorm.DB, error) {
	connectionString := config.TEST_DB_CONNECTION_STRING
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func createTestDBBookController(db *gorm.DB) (echo.HandlerFunc, error) {
	// bikin test model
	m := model.NewGormBookModel(db)
	// bikin controller
	c := CreateBookController(m)
	return c, nil
}

func testController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []model.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("expected one book, got: %#v", books)
		return
	}
	if books[0].Title != "Harry Potter" {
		t.Errorf("expected Harry Potter, got: %#v", books[0].Title)
	}
}

func TestDBBookController(t *testing.T) {
	// bikin db
	db, err := createTestDB()
	if err != nil {
		t.Error(err)
	}
	// bikin controller
	bookController, err := createTestDBBookController(db)
	if err != nil {
		t.Error(err)
	}
	// insert data baru
	db.Save(&model.Book{Title: "Harry Potter"})
	// test controller
	testController(t, bookController)
}

func TestMockBookController(t *testing.T) {
	m := model.NewMockBookModel()
	bookController := CreateBookController(m)
	// insert data baru
	m.Insert(model.Book{Title: "Harry Potter"})
	// test controller
	testController(t, bookController)
}
