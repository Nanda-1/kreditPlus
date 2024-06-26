package repositories

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
)

func CreateLoan(loan *models.Loan, limit int) error {
	return database.DB.Create(loan).Error
}

func GetLoanByID(id uint) (models.Loan, error) {
	var loan models.Loan
	err := database.DB.First(&loan, id).Error
	return loan, err
}

func UpdateLoan(loan *models.Loan) error {
	return database.DB.Save(loan).Error
}

func DeleteLoan(id uint) error {
	return database.DB.Delete(&models.Loan{}, id).Error
}
