package transport

import (
	"TIPPr4/internal/controllers"
	"TIPPr4/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users:user_id", controllers.GetUserById())
	incomingRoutes.GET("/users", controllers.GetPaginatedUsers())
	incomingRoutes.GET("/users", controllers.GetAllUsers())
	incomingRoutes.PATCH("/users", controllers.UpdateUser())
}
