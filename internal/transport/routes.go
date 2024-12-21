package transport

import (
	"TIPPr4/internal/controllers/cartControllers"
	"TIPPr4/internal/controllers/gameControllers"
	"TIPPr4/internal/controllers/orderControllers"
	"TIPPr4/internal/controllers/userControllers"
	"TIPPr4/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// Группировка маршрутов для аутентификации
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", userControllers.Signup()) // Регистрация пользователя
			auth.POST("/login", userControllers.Login())   // Логин пользователя
		}

		users := v1.Group("/users")
		{
			users.Use(middleware.Authenticate())
			users.GET("/:user_id", userControllers.GetUserById())        // Получение пользователя по ID
			users.GET("/paginated", userControllers.GetPaginatedUsers()) // Получение списка пользователей с пагинацией
			users.GET("/", userControllers.GetAllUsers())                // Получение всех пользователей (можно сделать приватным, например, для админов)
			users.PATCH("/:user_id", userControllers.UpdateUser())       // Обновление данных пользователя
		}

		roles := v1.Group("/roles")
		{
			roles.Use(middleware.Authenticate())
			roles.POST("/", userControllers.CreateRole())
			roles.GET("/", userControllers.GetAllRoles())
			roles.GET("/:role_id", userControllers.GetRoleById())
			roles.PATCH("/:role_id", userControllers.UpdateRole())
			roles.DELETE("/:role_id", userControllers.DeleteRole())
		}
		categories := v1.Group("/categories")
		{
			categories.Use(middleware.Authenticate())
			categories.POST("/", gameControllers.CreateCategory())
			categories.GET("/", gameControllers.GetAllCategories())
			categories.GET("/paginated", gameControllers.GetPaginatedCategories())
			categories.GET("/:category_id", gameControllers.GetCategory())
			categories.PATCH("/:category_id", gameControllers.UpdateCategory())
			categories.DELETE("/:category_id", gameControllers.DeleteCategory())
		}
		developers := v1.Group("/developers")
		{
			developers.Use(middleware.Authenticate())
			developers.POST("/", gameControllers.CreateDeveloper())
			developers.GET("/", gameControllers.GetAllDevelopers())
			developers.GET("/paginated", gameControllers.GetPaginatedDevelopers())
			developers.GET("/:developer_id", gameControllers.GetDeveloper())
			developers.PATCH("/:developer_id", gameControllers.UpdateDeveloper())
			developers.DELETE("/:developer_id", gameControllers.DeleteDeveloper())
		}
		restricts := v1.Group("/restricts")
		{
			restricts.Use(middleware.Authenticate())
			restricts.POST("/", gameControllers.CreateRestrict())
			restricts.GET("/", gameControllers.GetAllRestricts())
			restricts.GET("/paginated", gameControllers.GetPaginatedRestricts())
			restricts.GET("/:restrict_id", gameControllers.GetRestrict())
			restricts.PATCH("/:restrict_id", gameControllers.UpdateRestrict())
			restricts.DELETE("/:restrict_id", gameControllers.DeleteRestrict())
		}
		games := v1.Group("/games")
		{
			games.Use(middleware.Authenticate())
			games.POST("/", gameControllers.CreateGame())
			games.GET("/", gameControllers.GetAllGames())
			games.GET("/paginated", gameControllers.GetPaginatedGames())
			games.GET("/:game_id", gameControllers.GetGame())
			games.PATCH("/:game_id", gameControllers.UpdateGame())
			games.DELETE("/:game_id", gameControllers.DeleteGame())
		}
		cart := v1.Group("/cart")
		{
			cart.Use(middleware.Authenticate())
			cart.POST("/:user_id", cartControllers.AddGameToCart())
			cart.GET("/:user_id", cartControllers.GetCart())
			cart.DELETE("/:user_id/remove/:game_id", cartControllers.RemoveGameFromCart())
			cart.DELETE("/:user_id/clear", cartControllers.ClearCart())
		}
		order := v1.Group("/order")
		{
			order.Use(middleware.Authenticate())
			order.POST("/:user_id/create", orderControllers.CreateOrderFromCart())
			order.GET("/:order_id", orderControllers.GetOrderByID())
			order.GET("/user/:user_id", orderControllers.GetUserOrders())
			order.GET("/", orderControllers.GetAllOrders())
		}
	}

}
