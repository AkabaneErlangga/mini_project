package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func InitRoute(e *echo.Echo) {

	//e.POST("/login", controller.LoginController)
	//e.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRETKEY"))))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRETKEY")),
		Skipper: func(c echo.Context) bool {
			// Skip middleware if path is equal 'login'
			if c.Request().URL.Path == "/user/login" {
				return true
			}
			return false
		},
	}))
	setCategoryRoute(e)
	setUserRoute(e)
	setEmployeeRoute(e)
	setCategoryRoute(e)
	setItemRoute(e)
	setOrderRoute(e)
}