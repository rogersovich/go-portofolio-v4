package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectTechnology struct {
	ID           uint `gorm:"primaryKey" `
	ProjectID    int  `gorm:"column:project_id"`
	TechnologyID int  `gorm:"column:technology_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index" `
}
