package controller

import (
	"github.com/labstack/echo/v4"
	"main/model"
	"main/services"
	"net/http"
	"strconv"
	"time"
)

func GetAllEmployees(c echo.Context) error {
	name := string('%')+c.QueryParam("name")+string('%')
	employees := services.GetAllEmployees(name)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get employee",
		"employees":   employees,
	})
}

func GetEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := services.GetEmployeeById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get employee",
		"employee":   employee,
	})
}

func AddEmployee(c echo.Context) error {
	employee := model.Employee{}
	c.Bind(&employee)
	layoutFormat := "2006-01-02 15:04:05"
	employee.HireDate, _ = time.Parse(layoutFormat, c.FormValue("hireDate"))

	employeeCreated, err := services.CreateEmployee(employee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new employee",
		"employee":    employeeCreated,
	})
}
