package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
}
