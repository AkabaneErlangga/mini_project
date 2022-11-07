package router

import (
	"github.com/labstack/echo/v4"
	"main/controller"
)

func setUserRoute(e *echo.Echo)  {
	user := e.Group("/user")
	{
		user.GET("", controller.GetAllUsers)
		user.GET("/:id", controller.GetUser)
		user.POST("", controller.AddUser)
		user.PUT("/:id", controller.UpdateUser)
		user.DELETE("/:id", controller.DeleteUser)
		user.POST("/login", controller.LoginController)

	}
}