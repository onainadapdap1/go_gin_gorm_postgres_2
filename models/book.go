package models

import "time"

type Book struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"not null;type:varchar(191)" binding:"required"`
	Author string `gorm:"not null;type:varchar(191)" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}