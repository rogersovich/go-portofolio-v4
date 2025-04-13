package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectContentImage struct {
	ID            uint     `gorm:"primaryKey" `
	ProjectId     *int     `gorm:"type:int"`
	Project       *Project `gorm:"foreignKey:ProjectId"`
	ImageUrl      string   `gorm:"type:varchar(500)" `
	ImageFileName string   `gorm:"type:varchar(255)" `
	IsUsed        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" `
}
