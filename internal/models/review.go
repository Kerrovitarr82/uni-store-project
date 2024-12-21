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
	GameId      int            `json:"game_id"`
	Game        Game           `gorm:"foreignKey:GameId" json:"gameControllers"`
	UserID      int            `json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"userControllers"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
