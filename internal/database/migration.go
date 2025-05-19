package database

import (
	"log"
	"uniStore/internal/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Game{},
		&models.Developer{},
		&models.Category{},
		&models.Restrict{},
		&models.Review{},
		&models.Order{},
		&models.ShoppingCart{},
		&models.Favorite{},
		&models.Library{},
		&models.LibraryGame{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration succeeded!")
}
