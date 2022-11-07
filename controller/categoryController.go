package controller

import (
	"github.com/labstack/echo/v4"
	"main/config"
	"main/model"
	"main/services"
	"net/http"
	"strconv"
)

func GetAllCategories(c echo.Context) error {
	category := services.GetAllCategories()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get category",
		"categories":   category,
	})
}
func GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := services.GetCategoryById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get category",
		"category":   category,
	})
}

func AddCategory(c echo.Context) error {
	ctg := model.Category{}

	if err := c.Bind(&ctg); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	createdCategory, err := services.CreateCategory(ctg)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new category",
		"category":    createdCategory,
	})
}

func DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteCategory(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category",
	})
}

func UpdateCategory(c echo.Context) error {
	var ctg model.Category
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.First(&ctg, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctg.Name = c.FormValue("name")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update category",
		"items":   ctg,
	})
}