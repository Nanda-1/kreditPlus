package repositories

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
)

func CreateNotification(notification *models.Notification) error {
	return database.DB.Create(notification).Error
}

func GetNotificationByID(id uint) (models.Notification, error) {
	var notification models.Notification
	err := database.DB.First(&notification, id).Error
	return notification, err
}

func UpdateNotification(notification *models.Notification) error {
	return database.DB.Save(notification).Error
}

func DeleteNotification(id uint) error {
	return database.DB.Delete(&models.Notification{}, id).Error
}
