package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectTechnology struct {
	ID           uint `gorm:"primaryKey" `
	ProjectID    int
	TechnologyID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" `
}
