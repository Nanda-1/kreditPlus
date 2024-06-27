package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	LoanID         uint    `gorm:"primaryKey" json:"load_id,omitempty"`
	CustomerID     uint    `json:"customer_id,omitempty"`
	ContractNumber string  `json:"contract_number,omitempty"`
	OTR            int     `json:"otr,omitempty"`
	AdminFee       int     `json:"admin_fee,omitempty"`
	Installment    int     `json:"installment,omitempty"`
	Interest       int     `json:"interest,omitempty"`
	AssetName      string  `json:"asset_name,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
