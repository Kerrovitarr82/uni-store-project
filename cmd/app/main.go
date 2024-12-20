package main

import (
	"TIPPr4/internal/database"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.ConnectToDB()
}

func main() {
	database.MigrateDB()

}

//TODO сделать jwt, прописать пути и их логику
