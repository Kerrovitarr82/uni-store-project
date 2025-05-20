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

// @host		127.0.0.1:8080
func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	// Инициализация конфигурации
	initConfig()

	// Инициализация базы данных
	database.ConnectToDB()
	database.MigrateDB()
	database.CheckAdminAndRoles()

	// Инициализация маршрутов
	router := gin.Default()
	transport.InitRoutes(router)

	// Получение порта и ip из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1"
	}
	docs.SwaggerInfo.Host = ip + ":8080"
	if !myUtils.IsProd() {
		log.Printf("Swagger UI is available at: http://127.0.0.1:%s/swagger/index.html\n", port)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Запуск сервера
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server on port %s: %v", port, err)
	}

}
