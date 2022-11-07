package router

import (
	"github.com/labstack/echo/v4"
	"main/controller"
)

func setCategoryRoute(e *echo.Echo)  {
	ctg := e.Group("/category")
	{
		ctg.GET("", controller.GetAllCategories)
		ctg.GET("/:id", controller.GetCategory)
		ctg.POST("", controller.AddCategory)
		ctg.PUT("/:id", controller.UpdateCategory)
		ctg.DELETE("/:id", controller.DeleteCategory)

	}
}