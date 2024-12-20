package database

import (
	"TIPPr4/internal/models"
	"log"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Product{},
		&models.Manufacturer{},
		&models.Supplier{},
		&models.Category{},
		&models.Review{},
		&models.Order{},
		&models.Delivery{},
		&models.ShoppingCart{},
		&models.ProductInCart{},
		&models.ProductInFavorite{},
		&models.Favorite{},
		&models.ProductInOrder{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration succeeded!")
}
