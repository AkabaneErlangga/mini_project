package router

import (
	"github.com/labstack/echo/v4"
	"main/controller"
)

func setOrderRoute(e *echo.Echo)  {
	order := e.Group("/order")
	{
		order.GET("", controller.GetAllOrders)
		order.GET("/:id", controller.GetOrder)
		order.POST("", controller.AddOrder)

	}
}