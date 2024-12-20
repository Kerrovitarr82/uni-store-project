package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"not null" json:"name" validate:"required,min=2,max=100"`
	SecondName   string         `gorm:"not null" json:"second_name" validate:"required,min=2,max=100"`
	ThirdName    string         `json:"third_name"`
	Email        string         `gorm:"unique;not null" json:"email" validate:"required,email"`
	PhoneNumber  string         `gorm:"unique;not null" json:"phone_number" validate:"required"`
	Password     string         `gorm:"not null" json:"password" validate:"required,min=6"`
	PaymentInfo  string         `json:"payment_info"`
	RoleID       int            `gorm:"not null;default:1" json:"role_id" validate:"required"`
	Role         Role           `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
}

type Role struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Type        string         `gorm:"not null" json:"type"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
