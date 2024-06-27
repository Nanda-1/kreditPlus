package models

import (
	"time"

	"gorm.io/gorm"
)

type Tenor struct {
	TenorID    uint `gorm:"not null"`
	TenorMount int  `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
