package router

import (
	"github.com/labstack/echo/v4"
	"main/controller"
)

func setItemRoute(e *echo.Echo)  {
	emp := e.Group("/item")
	{
		emp.GET("", controller.GetAllItem)
		emp.GET("/:id", controller.GetItem)
		emp.POST("", controller.AddItem)
		emp.PUT("/:id", controller.UpdateItem)
		emp.DELETE("/:id", controller.DeleteItem)

	}
}