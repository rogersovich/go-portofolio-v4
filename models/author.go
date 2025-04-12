package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID             uint    `gorm:"primaryKey" `
	Name           string  `gorm:"type:varchar(255)" `
	AvatarUrl      string  `gorm:"type:varchar(500)" `
	AvatarFileName *string `gorm:"type:varchar(255)" `
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" `
}
