package models

import (
	"time"

	"gorm.io/gorm"
)

type Limits struct {
	LimitId    uint   `gorm:"primaryKey" json:"id"`
	LimitValue string `gorm:"type:decimal(18,2)" json:"value"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
