package models

import "time"

type Book struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name_book" gorm:"not null;type:varchar(191)" binding:"required"`
	Author string `json:"author" gorm:"not null;type:varchar(191)" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}