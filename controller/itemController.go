package controller

import (
	"github.com/labstack/echo/v4"
	"main/model"
	"main/services"
	"net/http"
	"strconv"
)

func GetAllItem(c echo.Context) error {
	items := services.GetAllItems()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"items":   items,
	})
}
func GetItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := services.GetItemById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item",
		"items":   item,
	})
}

func AddItem(c echo.Context) error {
	item := model.Item{}
	err := c.Bind(&item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdItem, err := services.CreateItem(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    createdItem,
	})
}

func DeleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteItem(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

func UpdateItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := services.GetItemById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item.Brand = c.FormValue("brand")
	item.Type =  c.FormValue("type")
	item.Price, _ = strconv.Atoi(c.FormValue("price"))
	item.Stock, _ = strconv.Atoi(c.FormValue("stock"))
	item.CategoryId, _ =  strconv.Atoi(c.FormValue("categoryId"))
	item.LaunchYear =  c.FormValue("launchYear")
	item.Description = c.FormValue("description")

	createdItem, err := services.CreateItem(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//config.DB.Save(&item)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
		"items":   createdItem,
	})
}