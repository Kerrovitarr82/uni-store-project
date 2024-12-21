package cartControllers

import (
	"TIPPr4/internal/database"
	"TIPPr4/internal/helpers"
	"TIPPr4/internal/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// AddGameToCart godoc
// @Summary Add Game to Shopping Cart
// @Description Add a specific game to the user's shopping cart
// @Tags ShoppingCart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/cart/{user_id}/add/{game_id} [post]
func AddGameToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")
		gameID := c.Param("game_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Проверка наличия пользователя
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Проверка наличия игры
		var game models.Game
		if err := database.DB.WithContext(ctx).First(&game, gameID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Получение корзины пользователя
		var cart models.ShoppingCart
		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Добавление игры в корзину
		if err := database.DB.WithContext(ctx).Model(&cart).Association("Games").Append(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Game added to cart"})
	}
}

// RemoveGameFromCart godoc
// @Summary Remove Game from Shopping Cart
// @Description Remove a specific game from the user's shopping cart
// @Tags ShoppingCart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param game_id path int true "Game ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/cart/{user_id}/remove/{game_id} [delete]
func RemoveGameFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")
		gameID := c.Param("game_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// Проверка наличия пользователя
		var user models.User
		if err := database.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Проверка наличия игры
		var game models.Game
		if err := database.DB.WithContext(ctx).First(&game, gameID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Получение корзины пользователя
		var cart models.ShoppingCart
		if err := database.DB.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Удаление игры из корзины
		if err := database.DB.WithContext(ctx).Model(&cart).Association("Games").Delete(&game); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Game removed from cart"})
	}
}

// GetCart godoc
// @Summary Get Shopping Cart
// @Description Get the user's shopping cart contents
// @Tags ShoppingCart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.ShoppingCart
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/cart/{user_id} [get]
func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var cart models.ShoppingCart
		if err := database.DB.WithContext(ctx).Preload("Games").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, cart)
	}
}

// ClearCart godoc
// @Summary Clear Shopping Cart
// @Description Clear all games from the user's shopping cart
// @Tags ShoppingCart
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/cart/{user_id}/clear [delete]
func ClearCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		userID := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var cart models.ShoppingCart
		if err := database.DB.WithContext(ctx).Preload("Games").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Shopping cart not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := database.DB.WithContext(ctx).Model(&cart).Association("Games").Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Shopping cart cleared"})
	}
}
