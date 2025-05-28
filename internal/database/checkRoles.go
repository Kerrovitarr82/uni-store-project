package database

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
	"uniStore/internal/models"
	"uniStore/internal/myUtils"
)

func CheckRoles() {
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
	adminEmail := os.Getenv("FIRST_ADMIN_EMAIL")

	if err := DB.WithContext(ctx).First(&user, "email = ?", adminEmail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPass, err := myUtils.HashPassword(os.Getenv("FIRST_ADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Failed to hash password for admin")
				return
			}
			user = models.User{
				Name:        "Admin",
				SecondName:  "1",
				Email:       adminEmail,
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
