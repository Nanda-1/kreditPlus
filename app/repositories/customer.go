package repositories

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
)

func StoreCostumer(costumer *models.Customer) (*models.Customer, error) {
	result := database.DB.Create(costumer)
	if result.Error != nil {
		return nil, result.Error
	}
	return costumer, nil
}
func GetAllUsers() ([]models.Customer, error) {
	var users []models.Customer
	err := database.DB.Find(&users).Error
	return users, err
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
