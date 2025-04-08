package models

import (
	"time"

	"gorm.io/gorm"
)

type Technology struct {
	ID              uint           `gorm:"primaryKey" faker:"-"`
	Name            string         `gorm:"type:varchar(100);not null" faker:"-"`
	LogoURL         string         `gorm:"type:varchar(500)" faker:"-"`
	DescriptionHTML string         `gorm:"type:text" faker:"sentence"`
	IsMajor         bool           `gorm:"column:is_major" faker:"-"`
	CreatedAt       time.Time      `faker:"-"`
	UpdatedAt       time.Time      `faker:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" faker:"-"`
}
