package router

import (
	"github.com/labstack/echo/v4"
	"main/controller"
)

func setEmployeeRoute(e *echo.Echo)  {
	emp := e.Group("/employee")
	{
		emp.GET("", controller.GetAllEmployees)
		emp.GET("/:id", controller.GetEmployee)
		emp.POST("", controller.AddEmployee)

	}
}