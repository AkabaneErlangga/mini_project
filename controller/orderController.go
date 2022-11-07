package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"main/model"
	"main/services"
	"net/http"
	"strconv"
	"time"
)

func GetAllOrders(c echo.Context) error {
	order := services.GetAllOrders()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get order",
		"orders":  order,
	})
}

func GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := services.GetOrderById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get order",
		"orders":  order,
	})
}

func AddOrder(c echo.Context) error {
	var order model.Order
	err := c.Bind(&order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order.Date = time.Now()
	order.TransactionId = uuid.New().String()

	createdOrder, err := services.CreateOrder(order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get order",
		"orders":  createdOrder,
	})
}
