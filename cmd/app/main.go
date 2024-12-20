package main

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/transport"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// @title		Game Store
// @version		1.0
// @description	REST-API for game store

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	// Инициализация конфигурации
	initConfig()

	// Инициализация базы данных
	database.ConnectToDB()
	database.MigrateDB()

	// Инициализация маршрутов
	router := gin.Default()
	transport.InitRoutes(router)

	// Получение порта из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Запуск сервера
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}
}
