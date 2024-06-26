package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	NotifID   uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
