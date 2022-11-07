package services

import (
	"main/config"
	"main/model"
)

func GetAllItems() []model.Item {
	var items []model.Item

	if err := config.DB.Find(&items).Error; err != nil {
		return []model.Item{}
	}
	return items
}

func GetItemById(id int) (model.Item, error) {
	var item model.Item

	if err := config.DB.First(&item, id).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil

}

func CreateItem(item model.Item) (model.Item, error) {
	if err := config.DB.Save(&item).Error; err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func DeleteItem(id int) error {
	var item model.Item

	if err := config.DB.Delete(&item,id).Error; err != nil {
		return err
	}
	return nil
}
