package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectFeature struct {
	ID            uint    `gorm:"primaryKey" `
	Description   string  `gorm:"type:text" `
	ImageUrl      string  `gorm:"type:varchar(500)" `
	ImageFileName *string `gorm:"type:varchar(255)" `
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" `
}
