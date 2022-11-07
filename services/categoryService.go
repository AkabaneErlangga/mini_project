package services

import (
	"main/config"
	"main/model"
)

func GetAllCategories() []model.Category {

	var category []model.Category

	if err := config.DB.Find(&category).Error; err != nil {
		return []model.Category{}
	}
	return category
}

func GetCategoryById(id int) (model.Category, error) {
	var category model.Category

	if err := config.DB.Model(&model.Category{}).Preload("Items").First(&category, id).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func CreateCategory(category model.Category) (model.Category, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil

}

func DeleteCategory(id int) error {
	var ctg model.Category
	if err := config.DB.Delete(&ctg,id).Error; err != nil {
		return err
	}
	return nil
}