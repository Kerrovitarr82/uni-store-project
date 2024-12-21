package middleware

import (
	"TIPPr4/internal/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из cookie с именем "token"
		clientToken, err := c.Cookie("token")
		clientRefreshToken, err := c.Cookie("refreshToken")
		if err != nil || clientToken == "" || clientRefreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing or invalid"})
			c.Abort()
			return
		}

		// Валидируем токен
		claims, err := helpers.ValidateToken(clientToken, clientRefreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Устанавливаем данные из токена в контекст
		c.Set("email", claims.Email)
		c.Set("first_name", claims.Name)
		c.Set("last_name", claims.SecondName)
		c.Set("uid", claims.Uid)
		c.Set("user_role", claims.UserRole)
		c.Next()
	}
}
