package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Rating      int            `gorm:"not null" json:"rating"`
	ProductID   int            `json:"product_id"`
	Product     Product        `gorm:"foreignKey:ProductID" json:"product"`
	UserID      int            `json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
