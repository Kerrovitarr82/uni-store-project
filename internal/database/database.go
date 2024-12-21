package database

import (
	"TIPPr4/internal/models"
	"TIPPr4/internal/myUtils"
	"context"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	CheckAdminAndRoles()
}

func CheckAdminAndRoles() {
	var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var role models.Role
	var user models.User

	if err := DB.WithContext(ctx).First(&role, "type = ?", "ADMIN").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = models.Role{
				Type:        "ADMIN",
				Description: "Позволяет проводить все операции",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			DB.Create(&role)
		}
	}
	if err := DB.WithContext(ctx).First(&role, "type = ?", "USER").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			role = models.Role{
				Type:        "USER",
				Description: "Позволяет проводить базовые операции юзера",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			DB.Create(&role)
		}
	}

	if err := DB.WithContext(ctx).First(&user, "role_id = ?", 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPass, err := myUtils.HashPassword(os.Getenv("FIRST_ADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Failed to hash password for admin")
				return
			}
			user = models.User{
				Name:        "Admin",
				SecondName:  "1",
				Email:       os.Getenv("FIRST_ADMIN_EMAIL"),
				PhoneNumber: "none",
				Password:    hashedPass,
				RoleID:      1,
				Role: models.Role{
					ID: 1,
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			DB.Create(&user)
		}
	}
}
