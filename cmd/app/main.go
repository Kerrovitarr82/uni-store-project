package main

import (
	_ "TIPPr4/api/docs"
	"TIPPr4/internal/database"
	"TIPPr4/internal/transport"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Printf("Swagger UI is available at: http://localhost:%s/swagger/index.html\n", port)
	// Запуск сервера
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}

}
