package models

import (
	"time"

	"gorm.io/gorm"
)

type Technology struct {
	ID              uint           `gorm:"primaryKey;autoIncrement;->;" faker:"-"`
	Name            string         `gorm:"type:varchar(100);not null" faker:"-"`
	LogoURL         string         `gorm:"type:varchar(500)" faker:"-"`
	DescriptionHTML string         `gorm:"type:text" faker:"sentence"`
	IsMajor         bool           `gorm:"column:is_major" faker:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index" faker:"-"`
	CreatedAt       time.Time      `gorm:"column:created_at" faker:"-"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" faker:"-"`
}
