package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Delivery struct {
	ID           int             `gorm:"primaryKey" json:"id"`
	OrderID      int             `json:"order_id"`
	ArrivalDate  time.Time       `json:"arrival_date"`
	DeliveryType string          `gorm:"not null" json:"delivery_type"`
	Address      string          `gorm:"not null" json:"address"`
	Cost         decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"cost"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}
