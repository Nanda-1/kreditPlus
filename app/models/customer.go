package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerID uint      `gorm:"primaryKey"`
	NIK        string    `gorm:"unique;not null"`
	FullName   string    `gorm:"not null"`
	LegalName  string    `gorm:"not null"`
	BirthPlace string    `gorm:"not null"`
	BirthDate  time.Time `gorm:"not null"`
	Salary     float64   `gorm:"not null"`
	IDPhoto    string    `gorm:"not null"`
	Selfie     string    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Tenors     []Tenor        `gorm:"foreignKey:UserID"`
	Loans      []Loan         `gorm:"foreignKey:UserID"`
}

type CustomerTenor struct {
	ID         uint    `gorm:"primaryKey"`
	CustomerID uint    `gorm:"foreignKey:CustomerID"`
	TenorID    uint    `gorm:""`
	Limit      float64 `gorm:"foreignKey:Limit"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
