package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	UserID    int             `json:"user_id"`
	User      User            `gorm:"foreignKey:UserID" json:"user"`
	Games     []Game          `gorm:"many2many:order_games;" json:"games"`
	TotalCost decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"total_cost"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}
