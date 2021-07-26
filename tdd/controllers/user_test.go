package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

// ada 3 ekspektasi:
// - status code == 200
// - response nya harus JSON
// - dalam response harus ada key users
func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	GetUserController(c)
	expectedStatusCode := http.StatusOK
	actualStatusCode := rec.Code
	if actualStatusCode != expectedStatusCode {
		t.Errorf("expected: %d, actual: %d", expectedStatusCode, actualStatusCode)
	}
	actualBody := rec.Body.Bytes()
	actualBodyData := map[string]interface{}{}
	if err := json.Unmarshal(actualBody, &actualBodyData); err != nil {
		t.Error(err)
	}
	_, keyExist := actualBodyData["users"]
	if !keyExist {
		t.Errorf("users should be exist, actual: %#v", actualBodyData)
	}
}
