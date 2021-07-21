package controllers

import (
	"alta/training/lib/database"
	"alta/training/middlewares"
	"alta/training/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func LoginUsersControllers(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)
	users, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// ini dari JWT token yang dikirim via header
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	// kalau loggedInUserId tidak sama dengan id yang dari parameter, kembalikan response 401
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access, you can only see your own")
	}
	users, err := database.GetDetailUsers(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
