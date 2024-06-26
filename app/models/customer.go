package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerID uint    `gorm:"primaryKey" json:"customer_id"`
	NIK        string  `gorm:"unique;not null" json:"nik"`
	FullName   string  `gorm:"not null" json:"full_name"`
	LegalName  string  `gorm:"not null" json:"legal_name"`
	BirthPlace string  `gorm:"not null" json:"birth_place"`
	BirthDate  string  `gorm:"not null" json:"birth_date"`
	Salary     float64 `gorm:"not null" json:"salary"`
	IDPhoto    string  `gorm:"not null" json:"photo_ktp"`
	Selfie     string  `gorm:"not null" json:"photo_selfie"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	// Tenors     []Tenor        `gorm:"foreignKey:UserID"`
	// Loans      []Loan         `gorm:"foreignKey:UserID"`
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
