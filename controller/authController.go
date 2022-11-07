package controller

import (
	"github.com/labstack/echo/v4"
	"main/config"
	"main/middleware"
	"main/model"
	"net/http"
)

func LoginController(c echo.Context) error {
	var user model.User
	c.Bind(&user)
	var userData model.User
	if err := config.DB.Where("username = ?", user.Username).First(&userData).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	isPasswordValid := middleware.CheckPasswordHash(user.Password, userData.Password)
	if !isPasswordValid {
		return c.JSON(http.StatusBadRequest, "password is not valid")
	}

	token, err := middleware.CreateToken(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"token":   token,
	})
}