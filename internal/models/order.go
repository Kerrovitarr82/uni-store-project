package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         int             `gorm:"primaryKey" json:"id"`
	UserID     int             `json:"user_id"`
	User       User            `gorm:"foreignKey:UserID" json:"user"`
	DeliveryID int             `json:"delivery_id"`
	Delivery   Delivery        `gorm:"foreignKey:DeliveryID" json:"delivery"`
	Status     string          `gorm:"not null" json:"status"`
	TotalCost  decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"total_cost"`
	Comment    string          `json:"comment"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type ProductInOrder struct {
	OrderID   int     `json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID" json:"order"`
	ProductID int     `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  uint64  `json:"quantity"`
}
