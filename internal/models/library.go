package models

import (
	"gorm.io/gorm"
	"time"
)

type Library struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	UserID    int            `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Games     []Game         `gorm:"many2many:library_games;" json:"games"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

type LibraryGame struct {
	LibraryID uint    `gorm:"primaryKey" json:"library_id"`
	GameID    uint    `gorm:"primaryKey" json:"game_id"`
	TimeSpent float64 `json:"time_spent"`
}
