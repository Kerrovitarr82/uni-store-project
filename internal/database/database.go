package database

import (
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
	for count := 0; count < 5; count++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return // успешное подключение
		}
		log.Println("Failed to connect to database. Retrying...")
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Could not connect to the database after 5 attempts:", err)
}
