package repositories

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
)

func CreateUser(user *models.Customer) error {
	return database.DB.Create(user).Error
}

func GetUserByID(id uint) (models.Customer, error) {
	var user models.Customer
	err := database.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user *models.Customer) error {
	return database.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.Customer{}, id).Error
}
