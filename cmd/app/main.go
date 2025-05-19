package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"uniStore/api/docs"
	_ "uniStore/api/docs"
	"uniStore/internal/database"
	"uniStore/internal/myUtils"
	"uniStore/internal/transport"
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
func main() {
	// Инициализация конфигурации
	initConfig()

	// Инициализация базы данных
	database.ConnectToDB()
	database.MigrateDB()
	database.CheckAdminAndRoles()

	// Инициализация маршрутов
	router := gin.Default()
	transport.InitRoutes(router)

	// Получение порта из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if myUtils.IsProd() {
		docs.SwaggerInfo.Host = "62.60.249.181:8080"
	} else {
		docs.SwaggerInfo.Host = "localhost:8080"
		log.Printf("Swagger UI is available at: http://localhost:%s/swagger/index.html\n", port)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Запуск сервера
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}

}
