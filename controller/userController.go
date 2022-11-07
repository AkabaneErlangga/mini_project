package controller

import (
	"github.com/labstack/echo/v4"
	"main/middleware"
	"main/model"
	"main/services"
	"net/http"
	"strconv"
)

func GetAllUsers(c echo.Context) error {
	users := services.GetAllUsers()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   users,
	})
}
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   user,
	})
}

func AddUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	hash, err := middleware.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = hash

	err = services.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user.Username,
	})
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := services.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hash, err := middleware.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Password = hash
	user.Username = c.FormValue("username")

	err = services.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
