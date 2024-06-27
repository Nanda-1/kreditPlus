package repositories

import (
	"errors"
	"kreditPlus/app/database"
	"kreditPlus/app/models"
)

// CreateLoan creates a new loan while handling concurrency
func CreateLoan(loan *models.Loan, TenorID int) error {
	// Start a new transaction
	tx := database.DB.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	// Lock the row for update to prevent race conditions
	var limit models.CustomerTenor
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&limit, TenorID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the limit is sufficient
	if limit.Limit < loan.Amount {
		tx.Rollback()
		return errors.New("insufficient limit")
	}

	// Deduct the loan amount from the limit
	limit.Limit -= loan.Amount

	// Save the updated limit
	if err := tx.Save(&limit).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create the loan
	if err := tx.Create(loan).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
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
