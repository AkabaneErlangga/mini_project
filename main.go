package main

import (
	"github.com/labstack/echo/v4"
	"main/config"
	"main/router"
)

func main() {

	config.InitDB()
	config.InitialMigration()

	e := echo.New()
	router.InitRoute(e)

	e.Logger.Fatal(e.Start(":8000"))
}