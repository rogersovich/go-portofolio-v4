package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID            uint `gorm:"primaryKey" `
	StatisticId   uint
	Statistic     Statistic
	Title         string  `gorm:"type:varchar(255)" `
	Description   string  `gorm:"type:text"`
	ImageURL      *string `gorm:"type:varchar(500)" `
	ImageFileName *string `gorm:"type:varchar(255)" `
	RepositoryURL string  `gorm:"type:varchar(500)"`
	Summary       string  `gorm:"type:text"`
	Status        string
	PublishedAt   *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index" `
}
