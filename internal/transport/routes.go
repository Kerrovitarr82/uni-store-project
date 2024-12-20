package transport

import (
	"TIPPr4/internal/controllers"
	"TIPPr4/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// Группировка маршрутов для аутентификации
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", controllers.Signup()) // Регистрация пользователя
			auth.POST("/login", controllers.Login())   // Логин пользователя
		}

		// Группировка маршрутов для пользователей
		users := v1.Group("/users")
		{
			users.Use(middleware.Authenticate())              // Применяем аутентификацию ко всем маршрутам в этой группе
			users.GET("/:user_id", controllers.GetUserById()) // Получение пользователя по ID
			users.GET("/", controllers.GetPaginatedUsers())   // Получение списка пользователей с пагинацией
			users.GET("/all", controllers.GetAllUsers())      // Получение всех пользователей (можно сделать приватным, например, для админов)
			users.PATCH("/", controllers.UpdateUser())        // Обновление данных пользователя
		}
	}

}
