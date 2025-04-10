package models

import (
	"time"

	"gorm.io/gorm"
)

type About struct {
	ID              uint   `gorm:"primaryKey" `
	Title           string `gorm:"type:varchar(255)" `
	AvatarFileName  string `gorm:"type:varchar(255)" `
	AvatarUrl       string `gorm:"type:varchar(500)" `
	DescriptionHTML string `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index" `
}
