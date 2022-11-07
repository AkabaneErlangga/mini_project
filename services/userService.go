package services

import (
	"main/config"
	"main/model"
)

func GetAllUsers() []model.User {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return []model.User{}
	}
	return users
}

func GetUserById(id int) (model.User, error) {
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil

}

func CreateUser(user model.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	var user model.User

	if err := config.DB.Delete(&user,id).Error; err != nil {
		return err
	}
	return nil
}
