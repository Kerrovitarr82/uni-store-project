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
	count := 0
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil && count != 5 {
		count++
		log.Println("Failed to connect to database. Restart connecting...")
		time.Sleep(5 * time.Second)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

}
