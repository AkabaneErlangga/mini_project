package services

import (
	"main/config"
	"main/model"
)

func GetAllOrders() []model.Order {
	var orders []model.Order

	if err := config.DB.Find(&orders).Error; err != nil {
		return []model.Order{}
	}
	return orders
}

func GetOrderById(id int) (model.Order, error) {
	var order model.Order

	if err := config.DB.First(&order, id).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil

}

func CreateOrder(order model.Order) (model.Order, error) {
	if err := config.DB.Save(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}
