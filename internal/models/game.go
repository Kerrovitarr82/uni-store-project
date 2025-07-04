package models

import (
	"gorm.io/gorm"
	"time"
)

type Game struct {
	ID             int            `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	Description    string         `json:"description"`
	Size           float64        `json:"size"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	AgeRestriction string         `json:"age_restriction"`
	DeveloperId    int            `json:"developer_id"`
	Developer      Developer      `gorm:"foreignKey:DeveloperId" json:"developer"`
	Categories     []Category     `gorm:"many2many:game_categories;" json:"categories"`
	Restricts      []Restrict     `gorm:"foreignKey:GameID" json:"restricts"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

type Developer struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Email       string         `gorm:"not null" json:"email"`
	Description string         `gorm:"not null" json:"description"`
	Country     string         `gorm:"not null" json:"country"`
	Games       []Game         `gorm:"foreignKey:DeveloperId" json:"games"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

type Category struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Games       []Game         `gorm:"many2many:game_categories;" json:"games"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}

type Restrict struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	GameID int    `gorm:"not null" json:"game_id"`
	Region string `gorm:"not null" json:"region"`
}
