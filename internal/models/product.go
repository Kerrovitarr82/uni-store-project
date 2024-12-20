package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID             int             `gorm:"primaryKey" json:"id"`
	Name           string          `gorm:"not null" json:"name"`
	Quantity       int             `gorm:"not null" json:"quantity"`
	Description    string          `json:"description"`
	Price          decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"price"`
	ManufacturerID int             `json:"manufacturer_id"`
	Manufacturer   Manufacturer    `gorm:"foreignKey:ManufacturerID" json:"manufacturer"`
	SupplierID     int             `json:"supplier_id"`
	Supplier       Supplier        `gorm:"foreignKey:SupplierID" json:"supplier"`
	CategoryID     int             `json:"category_id"`
	Category       Category        `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
}

type Manufacturer struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"not null" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Category struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Supplier struct {
	ID             int            `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `gorm:"not null" json:"email"`
	PaymentAddress string         `json:"payment_address"`
	Contract       string         `json:"contract"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
