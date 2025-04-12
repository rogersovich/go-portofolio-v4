package models

import (
	"time"

	"gorm.io/gorm"
)

type Statistic struct {
	ID        uint `gorm:"primaryKey" `
	Likes     int  `gorm:"type:int" `
	Views     int  `gorm:"type:int" `
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" `
}
