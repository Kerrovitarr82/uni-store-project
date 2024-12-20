package models

import (
	"gorm.io/gorm"
	"time"
)

type ShoppingCart struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	CustomerID int            `json:"customer_id"`
	User       User           `gorm:"foreignKey:CustomerID" json:"user"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ProductInCart struct {
	ProductID      int          `gorm:"primaryKey" json:"product_id"`
	Product        Product      `gorm:"foreignKey:ProductID" json:"product"`
	ShoppingCartID int          `gorm:"primaryKey" json:"shopping_cart_id"`
	ShoppingCart   ShoppingCart `gorm:"foreignKey:ShoppingCartID" json:"shopping_cart"`
	Quantity       int          `gorm:"not null"   json:"quantity"`
	CreatedAt      time.Time    `json:"created_at"`
}
