package services

import (
	"main/config"
	"main/model"
)

func GetAllEmployees(name string) []model.Employee {
	var employees []model.Employee

	if err := config.DB.Where("name LIKE ?", name).Find(&employees).Error; err != nil {
		return []model.Employee{}
	}
	return employees
}

func GetEmployeeById(id int) (model.Employee, error) {
	var employee model.Employee

	if err := config.DB.First(&employee, id).Error; err != nil {
		return model.Employee{}, err
	}
	return employee, nil

}

func CreateEmployee(employee model.Employee) (model.Employee, error) {
	if err := config.DB.Save(&employee).Error; err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}